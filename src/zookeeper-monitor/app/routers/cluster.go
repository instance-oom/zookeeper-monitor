package routers

import (
	"zookeeper-monitor/app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/cluster", &controllers.ClusterController{}, "get:Home")
	beego.Router("/cluster/detail/:id", &controllers.ClusterController{}, "get:Detail")
	beego.Router("/cluster/add", &controllers.ClusterController{}, "get:Add;post:Add")
	beego.Router("/cluster/edit/:id", &controllers.ClusterController{}, "get:Edit;post:Edit")
	beego.Router("/cluster/delete", &controllers.ClusterController{}, "delete:Delete")
}
