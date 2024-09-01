package PricingStrategy

import "ParkingLot/Models/enums"

type PricingStrategy interface {
	CalculateParkingTicketPrice(vehicleType enums.VehicleType) (int, error)
}
