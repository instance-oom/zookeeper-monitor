package controllers

import (
	"encoding/json"
	"strconv"
	"time"

	"zookeeper-monitor/app/jobs"
	"zookeeper-monitor/app/models"
)

//ServerController : ServerController
type ServerController struct {
	BaseController
}

//Home page
func (s *ServerController) Home() {
	servers, err := models.GetAllServers()
	if err != nil {
		s.showError(err.Error())
	}
	s.Data["pageTitle"] = "All Servers"
	s.Data["servers"] = servers
	s.display()
}

//Detail page
func (s *ServerController) Detail() {
	id, _ := strconv.Atoi(s.Ctx.Input.Param(":id"))
	server, err := models.GetServerByID(id)
	if err != nil {
		s.showError(err.Error())
	}
	status, err := models.GetStatus(id)
	if err != nil {
		s.showError(err.Error())
	}

	s.Data["pageTitle"] = "Server Detail"
	s.Data["server"] = server
	if len(status) > 0 {
		s.Data["latestStatus"] = status[len(status)-1]
	} else {
		s.Data["latestStatus"] = new(models.Status)
	}
	s.Data["status"] = status
	s.display()
}

//AddPage server page
func (s *ServerController) AddPage() {
	clusters, err := models.GetAllClusters()
	if err != nil {
		s.showError(err.Error())
	}
	s.Data["clusters"] = clusters
	s.Data["pageTitle"] = "Add Server"
	s.display("server/add")
}

//Add server
func (s *ServerController) Add() {
	server := new(models.Server)
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &server)
	if err != nil {
		s.ajaxMsg(err.Error(), MSG_ERR)
	}
	if server.ClusterID == 0 {
		s.ajaxMsg("ClusterID cannot be null or empty", MSG_ERR)
	}
	if server.IP == "" {
		s.ajaxMsg("IP cannot be null or empty", MSG_ERR)
	}
	if server.Port == "" {
		s.ajaxMsg("Port cannot be null or empty", MSG_ERR)
	}
	if server.Name == "" {
		s.ajaxMsg("Name cannot be null or empty", MSG_ERR)
	}
	if models.IsServerExist(server.IP, server.Port, server.Name) {
		s.ajaxMsg("Server is exists", MSG_ERR)
	}
	server.InUser = "ly61"
	server.InDate = time.Now()
	server.EditUser = "ly61"
	server.EditDate = time.Now()
	_, err = models.AddServer(server)
	if err != nil {
		s.ajaxMsg(err, MSG_ERR)
	} else {
		jobs.ServerChanged()
		s.ajaxMsg(server.ClusterID, MSG_OK)
	}
}

//EditPage server page
func (s *ServerController) EditPage() {
	id, _ := strconv.Atoi(s.Ctx.Input.Param(":id"))
	server, err := models.GetServerByID(id)
	if err != nil {
		s.showError(err.Error())
	}
	clusters, err := models.GetAllClusters()
	if err != nil {
		s.showError(err.Error())
	}
	s.Data["clusters"] = clusters
	s.Data["server"] = server
	s.Data["pageTitle"] = "Edit Cluster"
	s.display("server/edit")
}

//Edit server post
func (s *ServerController) Edit() {
	server := new(models.Server)
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &server)
	if err != nil {
		s.ajaxMsg(err.Error(), MSG_ERR)
	}
	if server.Name == "" {
		s.ajaxMsg("Server Name cannot be null or empty.", MSG_ERR)
	}
	if models.IsServerExistForEdit(server.IP, server.Port, server.Name, server.ID) {
		s.ajaxMsg("Server is exists", MSG_ERR)
	}
	server.EditUser = "ly61"
	server.EditDate = time.Now()
	err = server.Update()
	if err != nil {
		s.ajaxMsg(err, MSG_ERR)
	} else {
		jobs.ServerChanged()
		s.ajaxMsg(server.ClusterID, MSG_OK)
	}
}

//Delete server
func (s *ServerController) Delete() {
	var server models.Server
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &server)
	if err != nil {
		s.ajaxMsg(err.Error(), MSG_ERR)
	}
	if server.ID == 0 {
		s.ajaxMsg("ServerID cannot be null or empty", MSG_ERR)
	}
	err = models.DelServerByServerID(server.ID)
	if err != nil {
		s.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		jobs.ServerChanged()
		s.ajaxMsg("", MSG_OK)
	}
}
