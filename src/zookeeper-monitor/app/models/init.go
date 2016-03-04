package models

import (
	"fmt"
	"net/url"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //mysql driver
)

//Init is for orm init
func Init() {
	dbhost := os.Getenv("zkmonitor_dbhost")
	if dbhost == "" {
		dbhost = beego.AppConfig.String("db.host")
	}
	dbport := os.Getenv("zkmonitor_dbport")
	if dbport == "" {
		dbport = beego.AppConfig.String("db.port")
	}
	dbuser := os.Getenv("zkmonitor_dbuser")
	if dbuser == "" {
		dbuser = beego.AppConfig.String("db.user")
	}
	dbpassword := os.Getenv("zkmonitor_dbpassword")
	if dbpassword == "" {
		dbpassword = beego.AppConfig.String("db.password")
	}
	dbname := os.Getenv("zkmonitor_dbname")
	if dbname == "" {
		dbname = beego.AppConfig.String("db.name")
	}
	if dbport == "" {
		dbport = "3306"
	}
	timezone := beego.AppConfig.String("db.timezone")
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	fmt.Println(dsn)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", dsn, maxIdle, maxConn)

	orm.RegisterModel(new(Status), new(Server), new(Cluster), new(Mail))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
