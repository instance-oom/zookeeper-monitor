package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"zookeeper-monitor/app/jobs"
	"zookeeper-monitor/app/models"
)

//ServerController : ServerController
type ServerController struct {
	BaseController
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
	s.Data["status"] = status
	s.display()
}

//Add server page
func (s *ServerController) Add() {
	if s.isPost() {
		server := new(models.Server)
		server.ClusterID, _ = s.GetInt("cluster_id")
		server.Name = strings.TrimSpace(s.GetString("server_name"))
		server.IP = strings.TrimSpace(s.GetString("server_ip"))
		server.Port = strings.TrimSpace(s.GetString("server_port"))
		server.Description = strings.TrimSpace(s.GetString("server_desc"))
		server.InUser = "ly61"
		server.InDate = time.Now()
		server.EditUser = "ly61"
		server.EditDate = time.Now()

		_, err := models.AddServer(server)
		if err != nil {
			s.showError(err.Error())
		} else {
			jobs.ServerChanged()
			s.redirect("/cluster/detail/" + fmt.Sprintf("%d", server.ClusterID))
		}
	}
	clusters, err := models.GetAllClusters()
	if err != nil {
		s.showError(err.Error())
	}
	s.Data["clusters"] = clusters
	s.Data["pageTitle"] = "Add Server"
	s.display()
}

//Edit server page
func (s *ServerController) Edit() {
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

	if s.isPost() {
		server.Name = strings.TrimSpace(s.GetString("server_name"))
		server.IP = strings.TrimSpace(s.GetString("server_ip"))
		server.Port = strings.TrimSpace(s.GetString("server_port"))
		server.Description = strings.TrimSpace(s.GetString("server_desc"))
		server.EditUser = "ly61"
		server.EditDate = time.Now()
		err := server.Update()
		if err != nil {
			s.showError(err.Error())
		} else {
			jobs.ServerChanged()
			s.redirect("/cluster/detail/" + fmt.Sprintf("%d", server.ClusterID))
		}
	}

	s.Data["pageTitle"] = "Edit Cluster"
	s.Data["server"] = server
	s.display()
}

//Delete server
func (s *ServerController) Delete() {
	var server models.Server
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &server)
	if err != nil {
		s.ajaxMsg(err.Error(), MSG_ERR)
	}
	err = models.DelServerByServerID(server.ID)
	if err != nil {
		s.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		jobs.ServerChanged()
		s.ajaxMsg("", MSG_OK)
	}
}
