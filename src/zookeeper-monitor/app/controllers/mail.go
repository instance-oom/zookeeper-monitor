package controllers

import (
	"encoding/json"
	"zookeeper-monitor/app/jobs"
	"zookeeper-monitor/app/models"
)

//MailController : inclue index
type MailController struct {
	BaseController
}

//Home page
func (m *MailController) Home() {
	mail, err := models.GetMail()
	if err != nil {
		m.showError(err.Error())
	}

	m.Data["mail"] = mail
	m.Data["pageTitle"] = "Alert Mail Config"
	m.display()
}

//Add mail [post]
func (m *MailController) Add() {
	mail := new(models.Mail)
	err := json.Unmarshal(m.Ctx.Input.RequestBody, &mail)
	if err != nil {
		m.ajaxMsg(err.Error(), MSG_ERR)
	}
	if mail.Address == "" {
		m.ajaxMsg("Address cannot be null or empty", MSG_ERR)
	}
	_, err = models.AddMail(mail)
	if err != nil {
		m.ajaxMsg(err, MSG_ERR)
	} else {
		jobs.MailChanged()
		m.ajaxMsg("", MSG_OK)
	}
}

//Edit mail [post]
func (m *MailController) Edit() {
	mail := new(models.Mail)
	err := json.Unmarshal(m.Ctx.Input.RequestBody, &mail)
	if err != nil {
		m.ajaxMsg(err.Error(), MSG_ERR)
	}
	if mail.Address == "" {
		m.ajaxMsg("Address cannot be null or empty", MSG_ERR)
	}
	err = mail.Update()
	if err != nil {
		m.ajaxMsg(err, MSG_ERR)
	} else {
		jobs.MailChanged()
		m.ajaxMsg("", MSG_OK)
	}
}
