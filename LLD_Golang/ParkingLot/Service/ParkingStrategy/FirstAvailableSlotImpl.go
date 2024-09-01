package ParkingStrategy

import (
	"ParkingLot/Models"
	"ParkingLot/Repo"
	"errors"
)

type FirstAvailableSlotImpl struct {
	Repo *Repo.RepoHandler
}

func NewFirstAvailableSlotImpl(repo *Repo.RepoHandler) *FirstAvailableSlotImpl {
	return &FirstAvailableSlotImpl{Repo: repo}
}

func (fasl *FirstAvailableSlotImpl) FindAvailableSlot(vehicle *Models.Vehicle, parkingLotID string) (*Models.Slot, error) {
	parkingLot, ok := fasl.Repo.ParkingLots[parkingLotID]
	if !ok {
		return nil, errors.New("parking lot not exist")
	}

	var availableSlot *Models.Slot
	for _, floor := range parkingLot.ListOfFloors {
		for _, slot := range floor.ListSlots {
			if !slot.IsOccupied && slot.Type == vehicle.Type {
				availableSlot = &slot
				break
			}
		}
	}

	if availableSlot == nil {
		return nil, errors.New("parking slot not available for given vehicle")
	}
	return availableSlot, nil
}
