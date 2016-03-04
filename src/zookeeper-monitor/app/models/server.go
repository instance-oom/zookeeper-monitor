package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Server is the server information
type Server struct {
	ID          int       `orm:"column(ID)"`
	ClusterID   int       `orm:"column(ClusterID)"`
	IP          string    `orm:"column(IP)"`
	Port        string    `orm:"column(Port)"`
	Name        string    `orm:"column(Name)"`
	Description string    `orm:"column(Description);null"`
	InUser      string    `orm:"column(InUser)"`
	InDate      time.Time `orm:"column(InDate);type(datetime)"`
	EditUser    string    `orm:"column(EditUser);null"`
	EditDate    time.Time `orm:"column(EditDate);type(datetime);null"`
	IsRunning   bool      `orm:"column(IsRunning);type(bool)"`
	Mode        string    `orm:"column(Mode);null"`
	Statuses    []*Status `orm:"-"`
}

//AddServer : Add Server
func AddServer(server *Server) (int64, error) {
	server.InDate = time.Now()
	return orm.NewOrm().Insert(server)
}

//GetServerByID : Get Server By ServerID
func GetServerByID(serverID int) (*Server, error) {
	server := Server{ID: serverID}
	err := orm.NewOrm().Read(&server)
	return &server, err
}

//GetServersByClusterID : Get Servers By ClusterID
func GetServersByClusterID(clusterID int) ([]*Server, error) {
	var servers []*Server
	qs := orm.NewOrm().QueryTable("server")
	_, err := qs.Filter("ClusterID", clusterID).All(&servers)
	return servers, err
}

//GetAllServers : Get All Servers
func GetAllServers() ([]*Server, error) {
	var servers []*Server
	qs := orm.NewOrm().QueryTable("server")
	_, err := qs.All(&servers)
	return servers, err
}

//Update : Update server info
func (s *Server) Update(fields ...string) error {
	s.EditDate = time.Now()
	if _, err := orm.NewOrm().Update(s, "Name", "Description", "IsRunning", "Mode", "EditUser", "EditDate"); err != nil {
		return err
	}
	return nil
}

//DelServerByServerID : Delete Server By ServerID
func DelServerByServerID(serverID int) error {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return err
	}
	_, err = o.QueryTable("server").Filter("ID", serverID).Delete()
	_, err = o.QueryTable("status").Filter("ServerID", serverID).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

//DelServersByClusterID : Delete Servers By ClusterID
func DelServersByClusterID(clusterID int) error {
	_, err := orm.NewOrm().QueryTable("server").Filter("ClusterID", clusterID).Delete()
	return err
}

//IsServerExist : Analyzing server is present by server name 、ip or port
func IsServerExist(ip, port, name string) bool {
	cond := orm.NewCondition()
	exist := orm.NewOrm().QueryTable("server").SetCond(cond.AndCond(cond.And("IP", ip).And("Port", port)).OrCond(cond.And("Name", name))).Exist()
	return exist
}

//IsServerExistForEdit : Analyzing server is present
func IsServerExistForEdit(ip, port, name string, id int) bool {
	cond := orm.NewCondition()
	exist := orm.NewOrm().QueryTable("server").SetCond(cond.AndCond(cond.AndCond(cond.And("IP", ip).And("Port", port)).OrCond(cond.And("Name", name))).AndCond(cond.AndNot("ID", id))).Exist()
	return exist
}
