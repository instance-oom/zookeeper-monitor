package controllers

import (
	"encoding/json"
	"strconv"
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

//AddPage cluster page
func (c *ClusterController) AddPage() {
	c.Data["pageTitle"] = "Add Cluster"
	c.display("cluster/add")
}

//Add cluster post
func (c *ClusterController) Add() {
	cluster := new(models.Cluster)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cluster)
	if err != nil {
		c.ajaxMsg(err.Error(), MSG_ERR)
	}
	if cluster.Name == "" {
		c.ajaxMsg("Name cannot be null or empty", MSG_ERR)
	}
	if models.IsClusterExist(cluster.Name) {
		c.ajaxMsg("Name is exists", MSG_ERR)
	}
	cluster.InUser = "ly61"
	cluster.InDate = time.Now()
	cluster.EditUser = "ly61"
	cluster.EditDate = time.Now()
	_, err = models.AddCluster(cluster)
	if err != nil {
		c.ajaxMsg(err, MSG_ERR)
	} else {
		c.ajaxMsg("", MSG_OK)
	}
}

//EditPage cluster page
func (c *ClusterController) EditPage() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	cluster, err := models.GetClusterByID(id)
	if err != nil {
		c.showError(err.Error())
	}

	c.Data["pageTitle"] = "Edit Cluster"
	c.Data["cluster"] = cluster
	c.display("cluster/edit")
}

//Edit cluster post
func (c *ClusterController) Edit() {
	cluster := new(models.Cluster)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cluster)
	if err != nil {
		c.ajaxMsg(err.Error(), MSG_ERR)
	}
	if cluster.Name == "" {
		c.ajaxMsg("Cluster Name cannot be null or empty.", MSG_ERR)
	}
	if models.IsClusterExistForEdit(cluster.Name, cluster.ID) {
		c.ajaxMsg("Name is exists", MSG_ERR)
	}
	cluster.EditUser = "ly61"
	cluster.EditDate = time.Now()
	err = cluster.Update()
	if err != nil {
		c.ajaxMsg(err, MSG_ERR)
	} else {
		c.ajaxMsg("", MSG_OK)
	}
}

//Delete cluster
func (c *ClusterController) Delete() {
	var cluster models.Cluster
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cluster)
	if err != nil {
		c.ajaxMsg(err.Error(), MSG_ERR)
	}
	if cluster.ID == 0 {
		c.ajaxMsg("ClusterID cannot be null or empty", MSG_ERR)
	}
	err = models.DelClusterByID(cluster.ID)
	if err != nil {
		c.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		c.ajaxMsg("", MSG_OK)
	}
}
