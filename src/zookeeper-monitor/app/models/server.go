package models

import (
	"fmt"
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
	Statuses    []*Status `orm:"-"`
}

//AddServer : Add Server
func AddServer(server *Server) (int64, error) {
	if server.ClusterID == 0 {
		return 0, fmt.Errorf("ClusterID cannot be null or empty")
	}
	if server.IP == "" {
		return 0, fmt.Errorf("IP cannot be null or empty")
	}
	if server.Port == "" {
		return 0, fmt.Errorf("Port cannot be null or empty")
	}
	if server.Name == "" {
		return 0, fmt.Errorf("Name cannot be null or empty")
	}
	server.InDate = time.Now()
	return orm.NewOrm().Insert(server)
}

//GetServerByID : Get Server By ServerID
func GetServerByID(serverID int) (*Server, error) {
	if serverID == 0 {
		return nil, fmt.Errorf("ServerID cannot be null or empty")
	}
	server := Server{ID: serverID}
	err := orm.NewOrm().Read(&server)
	return &server, err
}

//GetServersByClusterID : Get Servers By ClusterID
func GetServersByClusterID(clusterID int) ([]*Server, error) {
	if clusterID == 0 {
		return nil, fmt.Errorf("ClusterID cannot be null or empty")
	}
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
	if s.Name == "" {
		return fmt.Errorf("Server Name cannot be null or empty.")
	}
	if _, err := orm.NewOrm().Update(s, fields...); err != nil {
		return err
	}
	return nil
}

//DelServerByServerID : Delete Server By ServerID
func DelServerByServerID(serverID int) error {
	if serverID == 0 {
		return fmt.Errorf("ServerID cannot be null or empty")
	}
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
	if clusterID == 0 {
		return fmt.Errorf("ClusterID cannot be null or empty")
	}
	_, err := orm.NewOrm().QueryTable("server").Filter("ClusterID", clusterID).Delete()
	return err
}
