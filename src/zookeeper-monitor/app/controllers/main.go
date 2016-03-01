package controllers

import "github.com/astaxie/beego"

//MainController : inclue index
type MainController struct {
	beego.Controller
}

//Index page
func (m *MainController) Index() {
	m.Data["pageTitle"] = "Home"
	tplname := "main/index.html"
	m.Layout = "layout.html"
	m.TplName = tplname
}
