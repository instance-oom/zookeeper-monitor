package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

//Cluster is cluster information
type Cluster struct {
	ID          int       `orm:"column(ID)"`
	Name        string    `orm:"column(Name)"`
	Description string    `orm:"column(Description);null"`
	InUser      string    `orm:"column(InUser)"`
	InDate      time.Time `orm:"column(InDate);type(datetime)"`
	EditUser    string    `orm:"column(EditUser);null"`
	EditDate    time.Time `orm:"column(EditDate);type(datetime);null"`
	Servers     []*Server `orm:"-"`
}

//AddCluster : Add Cluster
func AddCluster(cluster *Cluster) (int64, error) {
	if cluster.Name == "" {
		return 0, fmt.Errorf("Name cannot be null or empty")
	}
	return orm.NewOrm().Insert(cluster)
}

//GetAllClusters : Get All Cluster
func GetAllClusters() ([]*Cluster, error) {
	var clusters []*Cluster
	qs := orm.NewOrm().QueryTable("cluster")
	_, err := qs.All(&clusters)
	return clusters, err
}

//GetClusterByID : Get Cluster By ClusterID
func GetClusterByID(clusterID int) (*Cluster, error) {
	if clusterID == 0 {
		return nil, fmt.Errorf("ServerID cannot be null or empty")
	}
	cluster := Cluster{ID: clusterID}
	err := orm.NewOrm().Read(&cluster)
	return &cluster, err
}

//Update : Update Cluster Info
func (c *Cluster) Update(fields ...string) error {
	if c.Name == "" {
		return fmt.Errorf("Cluster Name cannot be null or empty.")
	}
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

//DelClusterByID : Delete Cluster By ClusterID
func DelClusterByID(clusterID int) error {
	if clusterID == 0 {
		return fmt.Errorf("ClusterID cannot be null or empty")
	}
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return err
	}
	_, err = o.QueryTable("cluster").Filter("ID", clusterID).Delete()
	_, err = o.QueryTable("server").Filter("ClusterID", clusterID).Delete()
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}
