package routers

import (
	"github.com/astaxie/beego"

	"parking_services/controllers"
)

func init() {
	beego.Router("/health", &controllers.HealthCheck{})
}