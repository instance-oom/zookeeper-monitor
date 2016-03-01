package models

//Server is the server information
type Server struct {
	ID          int
	ClusterID   int
	IP          string
	Port        string
	Name        string
	Description string
	InUser      string
	InDate      int64
	EditUser    string
	EditDate    int64
}
