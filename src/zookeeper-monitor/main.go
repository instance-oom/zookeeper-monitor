package main

import (
	"net/http"
	"text/template"

	"zookeeper-monitor/app/jobs"
	"zookeeper-monitor/app/models"
	_ "zookeeper-monitor/app/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var log = logs.NewLogger(10000)

func init() {
	log.SetLogger("console", "")
}

func main() {
	//Handle 404 error
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})
	models.Init()
	jobs.Init()

	go jobs.Run()

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
