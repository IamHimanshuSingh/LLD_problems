package Service

import (
	"ParkingLot/Models"
	"ParkingLot/Models/enums"
	"ParkingLot/Repo"
)

type FloorService interface {
	AddFloor(floor *Models.Floor) error
	RemoveFloor(floor *Models.Floor) error
	GetAllFreeSlots(floorID string) ([]Models.Slot, error)
	GetAllFreeSlotsByVehicleType(floorID string, vehicleType enums.VehicleType) ([]Models.Slot, error)
	GetAllOccupiedSlots(floorID string) ([]Models.Slot, error)
}

type FloorServiceHandler struct {
	Repo *Repo.RepoHandler
}

func NewFloorServiceHandler(repo *Repo.RepoHandler) *FloorServiceHandler {
	return &FloorServiceHandler{Repo: repo}
}

func (fsh *FloorServiceHandler) AddFloor(floor *Models.Floor) error {
	fsh.Repo.ParkingLots[floor.ParkingLotID].ListOfFloors = append(fsh.Repo.ParkingLots[floor.ParkingLotID].ListOfFloors, floor)
	return fsh.Repo.AddFloor(floor)
}

func (fsh *FloorServiceHandler) RemoveFloor(floor *Models.Floor) error {
	return fsh.Repo.RemoveFloor(floor)
}

func (fsh *FloorServiceHandler) GetAllFreeSlots(floorID string) ([]*Models.Slot, error) {
	result := make([]*Models.Slot, 0)
	allSlots, err := fsh.Repo.GetAllSlots(floorID)
	if err != nil {
		return nil, err
	}
	for _, slot := range allSlots {
		if !slot.IsOccupied {
			result = append(result, slot)
		}
	}
	return result, nil
}
func (fsh *FloorServiceHandler) GetAllFreeSlotsByVehicleType(floorID string, vehicleType enums.VehicleType) ([]*Models.Slot, error) {
	result := make([]*Models.Slot, 0)
	allSlots, err := fsh.Repo.GetAllSlots(floorID)
	if err != nil {
		return nil, err
	}
	for _, slot := range allSlots {
		if !slot.IsOccupied && slot.Type == vehicleType {
			result = append(result, slot)
		}
	}
	return result, nil
}

func (fsh *FloorServiceHandler) GetAllOccupiedSlots(floorID string) ([]*Models.Slot, error) {
	result := make([]*Models.Slot, 0)
	allSlots, err := fsh.Repo.GetAllSlots(floorID)
	if err != nil {
		return nil, err
	}
	for _, slot := range allSlots {
		if slot.IsOccupied {
			result = append(result, slot)
		}
	}
	return result, nil
}
