package main

import (
	"ParkingLot/Models"
	"ParkingLot/Models/enums"
	"ParkingLot/Repo"
	"ParkingLot/Service"
	"fmt"
)

func main() {
	parkingLot := Models.NewParkingLot("bengaluru", "PL001")

	// repo can be a singleton struct
	repo := Repo.NewRepoHandler()
	vehicleSvc := Service.NewVehicleHandler(repo)
	slotsSvc := Service.NewSlotServiceHandler(vehicleSvc, repo)
	ticketSvc := Service.NewTicketHandler(repo, slotsSvc)
	parkingLotSvc := Service.NewParkingLotServiceHandler(repo, ticketSvc, slotsSvc)

	// Adding Parking Lot
	err := parkingLotSvc.AddParkingLot(parkingLot)
	if err != nil {
		fmt.Println(err)
	}
	err = parkingLotSvc.GetAllParkingLots()
	if err != nil {
		fmt.Println(err)
	}

	// Adding Floors
	floorSvc := Service.NewFloorServiceHandler(repo)
	floor1 := Models.NewFloor("F1", parkingLot.ID, make([]*Models.Slot, 0))
	//floor2 := Models.NewFloor("F2", "PL001", make([]*Models.Slot, 0))
	//floor3 := Models.NewFloor("F3", "PL001", make([]*Models.Slot, 0))
	err = floorSvc.AddFloor(floor1)
	if err != nil {
		fmt.Println(err)
	}
	//err = floorSvc.AddFloor(floor2)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = floorSvc.AddFloor(floor3)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// Adding Vehicles
	vehicle1 := Models.NewVehicle("KA-01-GH-2343", enums.FourWheeler, "Owner_001")
	vehicle2 := Models.NewVehicle("KA-01-MH-2344", enums.TwoWheeler, "Owner_002")
	vehicle3 := Models.NewVehicle("KA-01-KH-2345", enums.TwoWheeler, "Owner_002")
	err = vehicleSvc.AddVehicle(vehicle1)
	if err != nil {
		fmt.Println(err)
	}
	err = vehicleSvc.AddVehicle(vehicle2)
	if err != nil {
		fmt.Println(err)
	}
	err = vehicleSvc.AddVehicle(vehicle3)
	if err != nil {
		fmt.Println(err)
	}

	// Adding Slots

	slot1 := Models.NewSlot("SL1", enums.TwoWheeler, floor1.ID)
	err = slotsSvc.AddSlot(slot1)
	if err != nil {
		fmt.Println(err)
	}
	slot2 := Models.NewSlot("SL2", enums.FourWheeler, floor1.ID)
	err = slotsSvc.AddSlot(slot2)
	if err != nil {
		fmt.Println(err)
	}
	slot3 := Models.NewSlot("SL3", enums.FourWheeler, floor1.ID)
	err = slotsSvc.AddSlot(slot3)
	if err != nil {
		fmt.Println(err)
	}

	// Park a vehicle
	ticket1, err := parkingLotSvc.ParkVehicle(vehicle1, parkingLot.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Ticket: %+v\n", ticket1)

	ticket2, err := parkingLotSvc.ParkVehicle(vehicle2, parkingLot.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Ticket: %+v\n", ticket2)

	ticket3, err := parkingLotSvc.ParkVehicle(vehicle3, parkingLot.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Ticket: %+v\n", ticket3)

	// UnPark a vehicle
	fmt.Printf("Unparking vehicle: %+v\n", ticket1)
	err = parkingLotSvc.UnParkVehicle(ticket1)
	if err != nil {
		fmt.Println(err)
	}

	// Get Free Slots on a floor
	slotIds, err := floorSvc.GetAllFreeSlots(floor1.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Free Slots on floorID: %s : ", floor1.ID)
	for _, slotId := range slotIds {
		fmt.Printf("%s, ", slotId)
	}

	// Get Occupied Slots on a floor
	slotIds, err = floorSvc.GetAllOccupiedSlots(floor1.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Occupied Slots on floorID: %s : ", floor1.ID)
	for _, slotId := range slotIds {
		fmt.Printf("%s, ", slotId)
	}

}
