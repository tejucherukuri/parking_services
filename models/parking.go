package models

import (
	"time"
)

type ParkingDetails struct {
   Id             int    	`json:"id"`
   ParkingFloor   string 	`json:"floor"`
   Block          string 	`json:"block"`
   OccupancyType  string 	`json:"occupancy_type"`//occupancy_type can be two_wheeler or four_wheeler
   Units          int    	`json:"units"` //count of vehicles which can be accomodated in the block
   Priority       int     	`json:"priority"`//to generate each parking slot and assign it to the user in a proper way
   CreatedDate    time.Time  `json:"-"`
   LastModified   time.Time   `json:"-"`
}


type ParkingEntries struct {
	Id             int          `json:"id"`
	VehicleNo      string       `json:"vehicle_no"`
	VehicleType    string       `json:"vehicle_type"`
	EntryTime      time.Time 	`json:"-"`
	ExitTime  		time.Time  	`json:"-"`
	AllotedSlot 	int 		`json:"parking_details_id`
	Status   		string 		`json:"status"` //status can be "active","paid" or "inactive"
	IsActive 		bool    	`json:"is_active"` //during exit time set the field to 0
	CreatedDate    time.Time    `json:"-"`
	LastModified   time.Time     `json:"-"`
}