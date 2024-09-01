package ParkingStrategy

import "ParkingLot/Models"

type ParkingStrategy interface {
	FindAvailableSlot(vehicle *Models.Vehicle, parkingLotID string) (*Models.Slot, error)
}

// We can use factory to get strategy object
