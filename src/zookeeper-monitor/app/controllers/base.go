package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

//BaseController base
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

//Prepare filter
func (b *BaseController) Prepare() {
	controllerName, actionName := b.GetControllerAndAction()
	b.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	b.actionName = strings.ToLower(actionName)
	//b.auth()

	b.Data["siteName"] = beego.AppConfig.String("site.name")
	b.Data["curRoute"] = b.controllerName + "." + b.actionName
	b.Data["curController"] = b.controllerName
	b.Data["curAction"] = b.actionName
}

func (b *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = tpl[0] + ".html"
	} else {
		tplname = b.controllerName + "/" + b.actionName + ".html"
	}
	b.Layout = "layout.html"
	b.TplName = tplname
}

func (b *BaseController) redirect(url string) {
	b.Redirect(url, 302)
	b.StopRun()
}

func (b *BaseController) isPost() bool {
	return b.Ctx.Request.Method == "POST"
}

func (b *BaseController) showError(args ...string) {
	b.Data["error"] = args[0]
	redirect := b.Ctx.Request.Referer()
	if len(args) > 1 {
		redirect = args[1]
	}

	b.Data["redirect"] = redirect
	b.Data["pageTitle"] = "Error"
	b.display("error/error")
	b.Render()
	b.StopRun()
}

func (b *BaseController) jsonResult(out interface{}, status int) {
	b.Data["json"] = out
	b.Ctx.Output.SetStatus(status)
	b.ServeJSON()
	b.StopRun()
}

func (b *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["msg"] = msg
	status := 200
	if msgno != MSG_OK {
		status = 500
	}
	b.jsonResult(out, status)
}
