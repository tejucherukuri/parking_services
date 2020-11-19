package routers

import (
	"github.com/astaxie/beego"
	"parking_services/controllers"
)


func init() {
	beego.Router("/health", &controllers.HealthCheck{})
	beego.Router("/fetchallocatedunits", &controllers.ParkingAllotmentController{}, "get:GetParkingOccupancyDetails")
	beego.Router("/user/parkingslot",&controllers.ParkingAllotmentController{}, "get:GetUserParkingDetails")
	beego.Router("/user/parkingslot",&controllers.ParkingAllotmentController{}, "post:AllocateUserSlot")
}