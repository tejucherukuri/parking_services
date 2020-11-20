package controllers

import (
	"parking_services/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
)

type ParkingAllotmentController struct {
	beego.Controller
}


/* GET request details to fetch the parked vehicle slot details based on vehicle number */

func (this *ParkingAllotmentController) GetUserParkingDetails() {
	db := models.NewOrm()
	vehicleno := this.GetString("vehicle_no")
	if vehicleno == "" { /* if vehicle number entered is empty */	
		this.Abort("Please enter the vehicle number")
		return
	}
	qry := `select parking_floor,block,entry_time from parking_details pt join parking_entries pe on pt.id = pe.alloted_slot where pe.vehicle_no = ? and pe.is_active=true`
	var totalsRawSeter orm.RawSeter
	totalsRawSeter = db.Raw(qry, vehicleno)
	var result []orm.ParamsList
	totalsRawSeter.ValuesList(&result)
    if (result == nil)  {
       this.Data["json"] = &map[string]string{
		"msg": "Please enter the correct vehicle number to fetch the details",
	   }
	} else {
		this.Data["json"] = result
	}
	this.ServeJSON()
}

/* POST Request to allocate the slot to the user */
func (this *ParkingAllotmentController) AllocateUserSlot() {
	db := models.NewOrm()
	vehicleno := this.GetString("vehicle_no")
	vehicletype := this.GetString("vehicle_type")
	/* Fetching the query details based on priority details */
	qry := `select pt.Id,pt.parking_floor,pt.block, pt.occupancy_type ,pt.priority, max(pt.units),max(pt.units) - count(pe.alloted_slot) as SlotsRem from parking_details pt left join parking_entries pe on pt.id = pe.alloted_slot where pt.occupancy_type=?
             group by pt.Id,pt.parking_floor,pt.block,pt.occupancy_type,pt.priority
             order by pt.priority`
	var totalsRawSeter orm.RawSeter
	totalsRawSeter = db.Raw(qry, vehicletype)
	var result []orm.ParamsList
	totalsRawSeter.ValuesList(&result)
	var parkallotment int
	for _, l := range result {
		totalunits,_ := strconv.Atoi(l[5].(string))
		allocatedunits,_ := strconv.Atoi(l[6].(string))
		if (totalunits - allocatedunits > 0) {
		  parkallotment,_ = strconv.Atoi(l[0].(string))
		  break
		}
	}

	userslot := &models.ParkingEntries {
		VehicleNo: vehicleno,
		AllotedSlot: parkallotment,
		VehicleType: vehicletype,
		EntryTime: time.Now(),
		Status: "active",
		IsActive: true,
	}
	_, err := db.InsertB(userslot.)
	if (err != nil) {
        this.Data["json"] = &map[string]string{
			"msg": "Couldn't generate a parking slot for the user.",
		}
	} else {
		/*TODO: Return the user slot allocated instead of static content with status */
		this.Data["json"] = &map[string]string{
			"msg": "User Entry Successfully Created",
		}
	}
	this.ServeJSON()
}


/* GET Request to find the free allocated slots for two wheeler and four wheelers */
func (this *ParkingAllotmentController) GetParkingOccupancyDetails()  {
	db := models.NewOrm()
	qry := `select pt.id,pt.parking_floor,pt.block,pt.occupancy_type,max(pt.units),max(pt.units) - count(pe.alloted_slot) as rem from parking_details pt left join parking_entries pe on pt.id = pe.alloted_slot 
	         group by pt.id,pt.parking_floor,pt.block,pt.occupancy_type,pe.alloted_slot`
	var fetchRawSeter orm.RawSeter
	fetchRawSeter = db.Raw(qry)
	var result []orm.ParamsList
	fetchRawSeter.ValuesList(&result)
	 var response []map[string]string
		for _,i := range result {
			details := map[string]string {
				"Floor": i[1].(string),
				"Block": i[2].(string),
				"Occupancy_type": i[3].(string),
				"Total": i[4].(string),
				"FreeSlots": i[5].(string),
			}
			response = append(response, details)
		}
	this.Data["json"] = response
	this.ServeJSON()
}
