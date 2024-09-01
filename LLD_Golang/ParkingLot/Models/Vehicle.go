package Models

import "ParkingLot/Models/enums"

type Vehicle struct {
	ID      string `json:"id"`
	Type    enums.VehicleType
	OwnerID string `json:"ownerId"`
}

func NewVehicle(ID string, Type enums.VehicleType, ownerID string) *Vehicle {
	return &Vehicle{ID: ID, Type: Type, OwnerID: ownerID}
}
