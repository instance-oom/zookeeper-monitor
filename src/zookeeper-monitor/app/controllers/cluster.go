package controllers

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
	"zookeeper-monitor/app/models"
)

//ClusterController : inclue add
type ClusterController struct {
	BaseController
}

//Home page
func (c *ClusterController) Home() {
	clusters, err := models.GetAllClusters()
	if err != nil {
		c.showError(err.Error())
	}
	c.Data["pageTitle"] = "All Cluster"
	c.Data["clusters"] = clusters
	c.display()
}

//Detail page
func (c *ClusterController) Detail() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	cluster, err := models.GetClusterByID(id)
	servers, err := models.GetServersByClusterID(id)
	if err != nil {
		c.showError(err.Error())
	}
	c.Data["pageTitle"] = "All Servers"
	c.Data["cluster"] = cluster
	c.Data["servers"] = servers
	c.display()
}

//Add cluster page
func (c *ClusterController) Add() {
	if c.isPost() {
		cluster := new(models.Cluster)
		cluster.Name = strings.TrimSpace(c.GetString("cluster_name"))
		cluster.Description = strings.TrimSpace(c.GetString("cluster_desc"))
		cluster.InUser = "ly61"
		cluster.InDate = time.Now()
		cluster.EditUser = "ly61"
		cluster.EditDate = time.Now()

		_, err := models.AddCluster(cluster)
		if err != nil {
			c.showError(err.Error())
		} else {
			c.redirect("/cluster")
		}
	}

	c.Data["pageTitle"] = "Add Cluster"
	c.display()
}

//Edit cluster page
func (c *ClusterController) Edit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	cluster, err := models.GetClusterByID(id)
	if err != nil {
		c.showError(err.Error())
	}

	if c.isPost() {
		cluster.Name = strings.TrimSpace(c.GetString("cluster_name"))
		cluster.Description = strings.TrimSpace(c.GetString("cluster_desc"))
		cluster.EditUser = "ly61"
		cluster.EditDate = time.Now()
		err := cluster.Update()
		if err != nil {
			c.showError(err.Error())
		} else {
			c.redirect("/cluster")
		}
	}

	c.Data["pageTitle"] = "Edit Cluster"
	c.Data["cluster"] = cluster
	c.display()
}

//Delete cluster
func (c *ClusterController) Delete() {
	var cluster models.Cluster
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cluster)
	if err != nil {
		c.ajaxMsg(err.Error(), MSG_ERR)
	}
	err = models.DelClusterByID(cluster.ID)
	if err != nil {
		c.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		c.ajaxMsg("", MSG_OK)
	}
}
