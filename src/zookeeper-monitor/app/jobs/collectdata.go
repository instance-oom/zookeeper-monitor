package jobs

import (
	"strconv"
	"time"

	"zookeeper-monitor/app/common"
	"zookeeper-monitor/app/models"

	"github.com/astaxie/beego"
	"github.com/samuel/go-zookeeper/zk"
)

var servers []*models.Server

//ServerChanged : Get servers when delete or create server
func ServerChanged() {
	var err error
	servers, err = models.GetAllServers()
	if err != nil {
		common.SendMail(err.Error(), "GetAllServers")
		return
	}
}

//Run : Begin to collect zookeeper server status
func Run() {
	collectTime := 300
	if time := beego.AppConfig.String("collect.time"); time != "" {
		collectTime, _ = strconv.Atoi(time)
	}
	timer := time.NewTicker(time.Second * time.Duration(collectTime))
	ServerChanged()
	for {
		select {
		case <-timer.C:
			go func() {
				for _, server := range servers {
					stats, _ := zk.FLWSrvr([]string{server.IP + ":" + server.Port}, time.Second*5)
					for _, stat := range stats {
						if stat.Error != nil {
							common.SendMail(stat.Error.Error(), server.IP)
							server.IsRunning = false
						} else {
							server.IsRunning = true
							status := new(models.Status)
							status.ServerID = server.ID
							status.Version = stat.Version
							status.AvgLatency = stat.AvgLatency
							status.MaxLatency = stat.MaxLatency
							status.MinLatency = stat.MinLatency
							status.PacketsReceived = stat.Received
							status.PacketsSend = stat.Sent
							status.NumAliveConnections = stat.Connections
							status.OutstandingRequests = stat.Outstanding
							status.ServerState = stat.Mode.String()
							status.ZnodeCount = stat.NodeCount
							status.WatchCount = -1
							status.EphemeralsCount = -1
							status.ApproximateDataSize = -1
							status.OpenFileDescriptorCount = -1
							status.MaxFileDescriptorCount = -1
							status.Followers = -1
							status.SyncedFollowers = -1
							status.PendingSyncs = -1
							models.AddStatus(status)
						}
						server.Update()
					}
				}
			}()
		}
	}
}
