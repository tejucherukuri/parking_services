package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "parking_services/routers"
	"time"
)

func init() {
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		beego.AppConfig.String("mysql_user")+":"+
			beego.AppConfig.String("mysql_password")+"@"+
			beego.AppConfig.String("mysql_db_host")+"/"+
			beego.AppConfig.String("mysql_db")+"?charset=utf8&parseTime=True")
}

func main() {
	beego.Run(beego.AppConfig.String("app_httpport"))
}