package controllers

//MainController : inclue index
type MainController struct {
	BaseController
}

//Index page
func (m *MainController) Index() {
	m.Data["pageTitle"] = "Home"
	m.display()
}
