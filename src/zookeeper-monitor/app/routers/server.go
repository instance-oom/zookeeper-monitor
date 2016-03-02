package routers

import (
	"zookeeper-monitor/app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/server/detail/:id", &controllers.ServerController{}, "get:Detail")
	beego.Router("/server/add", &controllers.ServerController{}, "get:Add;post:Add")
	beego.Router("/server/edit/:id", &controllers.ServerController{}, "get:Edit;post:Edit")
	beego.Router("/server/delete", &controllers.ServerController{}, "delete:Delete")
}
