# parking_services
Parking Allotment Services

Database Schema : Mentioned in the /data folder with some sample test data

Below are my assumptions behing the design and implementations:

Considering an example:

Suppose Parking Area of a Shopping Mall has 2 floors of parking - Level 1 and Level 2.

Each floor has n number of blocks - for example, Level 1 has block A, block B, block C and block D .

And each block can accomodate to n units of vehicles ( ex: Level 1 - block A can accodomate 14 four-wheelers and 5 two-wheelers , similarly block B can accomodate , 15 two-wheelers )

Each block has a priority assigned in the database , so that parking slots are assigned and filled one after the other in a proper manner based on the priorities which can be assigned in the db.

When an user enters the parking slot counter, vehicle number is  and also vehicle_type data (two_wheeler / four_wheeler) is taken  , based on the vehicle_type data, a query search is performed and based on priority of slots , a slot is assigned to the user.


For further extensible way:

1. Added the entry time and exit time fields in the database , so that performance metrics like below can be known:
  - Calculate the parking charges based on the number of hours spent.
  - Left over parked vehicles after a time period.
  - Track the status of vehicle whether it is active / paid / inactive.


API Routes : 


API Route to check if the server is up without any database configuration:

Route API : http://localhost:4646/health

Response: 
        {
        "status": "OK"
        }
        
1.

Route API : http://localhost:4646/fetchallocatedunits

Route Description:  GET Request to find the free allocated slots for two wheeler and four wheelers

Sample Response: 

[
        {
        "Block": "blockA",
        "Floor": "level1",
        "Occupancy_type": "two_wheeler",
        "Remaining": "6",
        "Total": "10"
        },
        {
        "Block": "blockB",
        "Floor": "level1",
        "Occupancy_type": "four_wheeler",
        "Remaining": "14",
        "Total": "15"
        },
        {
        "Block": "blockC",
        "Floor": "level1",
        "Occupancy_type": "two_wheeler",
        "Remaining": "5",
        "Total": "5"
        },
        {
        "Block": "blockC",
        "Floor": "level1",
        "Occupancy_type": "four_wheeler",
        "Remaining": "8",
        "Total": "8"
        },
        {
        "Block": "blockA",
        "Floor": "level2",
        "Occupancy_type": "four_wheeler",
        "Remaining": "19",
        "Total": "20"
        }
]


2. 

Route API : http://localhost:4646/user/parkingslot

Param: vehicle_no

Route Description:  GET request details to fetch the parked vehicle slot details based on vehicle number ( if user forgets )

Scenario :
 A. if the vehicle_no entered is empty
    Response: Please enter the vehicle number

 B. if the vehicle_no entered is not empty and valid ( exists in our db)
    Response: 

    [
            [
            "level1",
            "blockA",
            "2020-11-19T14:16:44Z"
            ]
     ]
 C. if the vehicle_no entered is not empty and not existing in our db.

    Response: 
     
     {
        "msg": "Please enter the correct vehicle number to fetch the details"
     }




3. 

Route API: 

Params: vehicle_no , vehicle_type

Route description: POST Request to allocate the slot to the user. 


Commands :

> go build
> ./parking_services
















