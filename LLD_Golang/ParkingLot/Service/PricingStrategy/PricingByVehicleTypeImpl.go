package PricingStrategy

import (
	"ParkingLot/Models/enums"
	"errors"
)

type PricingByVehicleTypeImpl struct {
}

func (PricingByVehicleTypeImpl) CalculateParkingTicketPrice(vehicleType enums.VehicleType) (int, error) {
	switch vehicleType {
	case enums.TwoWheeler:
		return 50, nil
	case enums.FourWheeler:
		return 100, nil
	}

	return -1, errors.New("invalid VehicleType")
}
