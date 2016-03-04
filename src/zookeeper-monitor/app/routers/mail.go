package routers

import (
	"zookeeper-monitor/app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/mail", &controllers.MailController{}, "*:Home")
	beego.Router("/mail/add", &controllers.MailController{}, "post:Add")
	beego.Router("/mail/edit", &controllers.MailController{}, "post:Edit")
}
