package main

import (
	 "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "parking_services/routers"
	"github.com/astaxie/beego/plugins/cors"
	"parking_services/models"
	"time"
)


func init() {
	orm.DefaultTimeLoc = time.UTC
	/* MySQL Database Configuration  */
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		beego.AppConfig.String("mysql_user")+":"+
			beego.AppConfig.String("mysql_password")+"@"+
			beego.AppConfig.String("mysql_db_host")+"/"+
			beego.AppConfig.String("mysql_db")+"?charset=utf8&parseTime=True")
	orm.RegisterModel(new(models.ParkingDetails),new(models.ParkingEntries))
}

func main() {
	/* To avoid CORS Policy */
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Access-Control-Allow-Credentials"},
		AllowCredentials: true,
	}))
	beego.Run(beego.AppConfig.String("app_httpport"))
}