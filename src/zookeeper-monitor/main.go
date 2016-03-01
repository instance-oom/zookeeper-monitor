package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"zookeeper-monitor/app/controllers"
	"zookeeper-monitor/app/models"

	"github.com/astaxie/beego"
	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	//Handle 404 error
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})
	models.Init()

	go func() {
		fmt.Println("test")
		timer := time.NewTicker(time.Second * 60)
		for {
			select {
			case <-timer.C:
				go func() {
					fmt.Printf("[%s] Begin to get server stat.\n", time.Now().Format("2006-01-02 15:04:05"))
					stats, _ := zk.FLWSrvr([]string{"127.0.0.1:8481"}, time.Second*5)
					for _, stat := range stats {
						if stat.Error != nil {
							fmt.Println(stat.Error)
							continue
						}
						status := new(models.Status)
						status.ServerID = 1
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
						result, err := models.AddStatus(status)
						fmt.Printf("%d %+v\n", result, err)
					}
				}()
			}
		}
	}()

	beego.Router("/", &controllers.MainController{}, "*:Index")

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
