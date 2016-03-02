package controllers

import "zookeeper-monitor/app/models"

//MainController : inclue index
type MainController struct {
	BaseController
}

//Index page
func (m *MainController) Index() {
	clusters, err := models.GetAllClusters()
	if err != nil {
		m.showError(err.Error())
	}
	servers, err := models.GetAllServers()
	if err != nil {
		m.showError(err.Error())
	}

	m.Data["clusters"] = clusters
	m.Data["servers"] = servers
	m.Data["pageTitle"] = "Home"
	m.display()
}
