package controllers

import (
	"github.com/astaxie/beego"
)

type HealthCheck struct {
	beego.Controller
}

func (this *HealthCheck) Get() {
	this.Data["json"] = map[string]string{"status": "OK"}
	this.ServeJSON()
}