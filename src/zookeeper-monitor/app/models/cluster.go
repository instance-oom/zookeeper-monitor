package models

//Cluster is cluster information
type Cluster struct {
	ID          int
	Name        string
	Description string
	InUser      string
	InDate      int64
	EditUser    string
	EditDate    int64
}
