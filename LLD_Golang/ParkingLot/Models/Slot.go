package Models

import "ParkingLot/Models/enums"

type Slot struct {
	ID         string
	Type       enums.VehicleType
	FloorID    string
	IsOccupied bool
}

func NewSlot(ID string, Type enums.VehicleType, floorID string) *Slot {
	return &Slot{ID: ID, Type: Type, FloorID: floorID, IsOccupied: false}
}
