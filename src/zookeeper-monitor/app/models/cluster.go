package models

import (
	"strings"
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
	return orm.NewOrm().Insert(cluster)
}

//GetAllClusters : Get All Cluster
func GetAllClusters() ([]*Cluster, error) {
	var clusters []*Cluster
	qs := orm.NewOrm().QueryTable("cluster")
	_, err := qs.All(&clusters)
	for i := 0; i < len(clusters); i++ {
		for j := i + 1; j < len(clusters); j++ {
			if strings.Compare(clusters[j].Name, clusters[i].Name) == -1 {
				clusters[i], clusters[j] = clusters[j], clusters[i]
			}
		}
	}
	return clusters, err
}

//GetClusterByID : Get Cluster By ClusterID
func GetClusterByID(clusterID int) (*Cluster, error) {
	cluster := Cluster{ID: clusterID}
	err := orm.NewOrm().Read(&cluster)
	return &cluster, err
}

//Update : Update Cluster Info
func (c *Cluster) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, "Name", "Description", "EditUser", "EditDate"); err != nil {
		return err
	}
	return nil
}

//DelClusterByID : Delete Cluster By ClusterID
func DelClusterByID(clusterID int) error {

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

//IsClusterExist : Analyzing cluster is present by name
func IsClusterExist(name string) bool {
	exist := orm.NewOrm().QueryTable("cluster").Filter("Name", name).Exist()
	return exist
}

//IsClusterExistForEdit : Analyzing cluster is present
func IsClusterExistForEdit(name string, id int) bool {
	exist := orm.NewOrm().QueryTable("cluster").Filter("Name", name).Exclude("ID", id).Exist()
	return exist
}
