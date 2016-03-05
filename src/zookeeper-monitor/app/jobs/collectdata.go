package jobs

import (
	"os"
	"strconv"
	"time"

	"zookeeper-monitor/app/common"
	"zookeeper-monitor/app/models"

	"github.com/astaxie/beego"
	"github.com/samuel/go-zookeeper/zk"
)

var servers []*models.Server
var mail *models.Mail

//Init : init base data
func Init() {
	MailChanged()
	ServerChanged()
}

//ServerChanged : Get servers when delete or create server
func ServerChanged() {
	var err error
	servers, err = models.GetAllServers()
	if err != nil {
		common.SendMail(err.Error(), mail.Address, "GetAllServers")
		return
	}
}

//MailChanged : Get mail info
func MailChanged() {
	var err error
	mail, err = models.GetMail()
	if err != nil || mail == nil {
		common.SendMail(err.Error(), "lon.l.yang@newegg.com", "GetAlertMailError")
		return
	}
}

//Run : Begin to collect zookeeper server status
func Run() {
	collectTime := 300
	timefromenv := os.Getenv("zkmonitor_collecttime")
	if timefromenv != "" {
		collectTime, _ = strconv.Atoi(timefromenv)
	} else {
		if timefromconfig := beego.AppConfig.String("collect.time"); timefromconfig != "" {
			collectTime, _ = strconv.Atoi(timefromconfig)
		}
	}
	timeToUpdateServerStatus := time.NewTicker(time.Second * time.Duration(collectTime))
	timeToSaveStatus := time.NewTicker(time.Minute * 5)
	for {
		select {
		case <-timeToUpdateServerStatus.C:
			go func() {
				for _, server := range servers {
					stats, _ := zk.FLWSrvr([]string{server.IP + ":" + server.Port}, time.Second*5)
					for _, stat := range stats {
						if stat.Error != nil {
							common.SendMail(stat.Error.Error(), mail.Address, server.IP)
							server.IsRunning = false
							server.Mode = "Unknown"
						} else {
							server.IsRunning = true
							server.Mode = stat.Mode.String()
						}
						server.EditUser = "job"
						server.Update()
					}
				}
			}()
		case <-timeToSaveStatus.C:
			go func() {
				for _, server := range servers {
					stats, _ := zk.FLWSrvr([]string{server.IP + ":" + server.Port}, time.Second*5)
					for _, stat := range stats {
						if stat.Error == nil {
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
					}
				}
			}()
		}
	}
}
