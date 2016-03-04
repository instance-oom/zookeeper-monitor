package routers

import (
	"zookeeper-monitor/app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/server", &controllers.ServerController{}, "get:Home")
	beego.Router("/server/detail/:id", &controllers.ServerController{}, "get:Detail")
	beego.Router("/server/add", &controllers.ServerController{}, "get:AddPage;post:Add")
	beego.Router("/server/edit/:id", &controllers.ServerController{}, "get:EditPage;post:Edit")
	beego.Router("/server/delete", &controllers.ServerController{}, "delete:Delete")
}
