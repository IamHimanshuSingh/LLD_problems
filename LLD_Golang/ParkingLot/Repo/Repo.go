package Repo

import (
	"ParkingLot/Models"
	"errors"
	"fmt"
)

// Repo Currently keeping all the repo functions as part of the same interface,
// but in real world, each microservice can have its own DB
type Repo interface {
	AddParkingLot(parkingLot *Models.ParkingLot) error
	RemoveParkingLot(parkingLot *Models.ParkingLot) error
	GetAllParkingLots() error
	AddFloor(floor *Models.Floor) error
	RemoveFloor(floor *Models.Floor) error
	GetAllSlots(floorID string) ([]Models.Slot, error)
	AddSlot(slot *Models.Slot) error
	RemoveSlot(slot *Models.Slot) error
	AddTicket(ticket *Models.Ticket) error
	AddVehicle(vehicle *Models.Vehicle) error
}

type RepoHandler struct {
	ParkingLots map[string]*Models.ParkingLot
	Floors      map[string]*Models.Floor
	Slots       map[string]*Models.Slot
	Tickets     map[string]*Models.Ticket
	Vehicles    map[string]*Models.Vehicle
}

// NewRepoHandler TODO: make this singleton
func NewRepoHandler() *RepoHandler {
	return &RepoHandler{
		ParkingLots: make(map[string]*Models.ParkingLot),
		Floors:      make(map[string]*Models.Floor),
		Slots:       make(map[string]*Models.Slot),
		Tickets:     make(map[string]*Models.Ticket),
		Vehicles:    make(map[string]*Models.Vehicle)}
}

func (rh *RepoHandler) AddParkingLot(parkingLot *Models.ParkingLot) error {
	if parkingLot == nil {
		return errors.New("ParkingLot is nil")
	}
	rh.ParkingLots[parkingLot.ID] = parkingLot
	return nil
}

func (rh *RepoHandler) RemoveParkingLot(parkingLot *Models.ParkingLot) error {
	if parkingLot == nil {
		return errors.New("ParkingLot is nil")
	}
	delete(rh.ParkingLots, parkingLot.ID)
	return nil
}

func (rh *RepoHandler) GetAllParkingLots() error {
	if rh.ParkingLots == nil {
		return errors.New("GetAllParkingLots is nil")
	}
	for _, parkingLot := range rh.ParkingLots {
		fmt.Printf("ParkingLotID: %s\n", parkingLot.ID)
	}
	return nil
}

func (rh *RepoHandler) AddFloor(floor *Models.Floor) error {
	if rh.Floors == nil {
		rh.Floors = make(map[string]*Models.Floor)
	}
	rh.Floors[floor.ID] = floor
	return nil
}

func (rh *RepoHandler) RemoveFloor(floor *Models.Floor) error {
	if floor == nil {
		rh.Floors = make(map[string]*Models.Floor)
	}
	delete(rh.Floors, floor.ID)
	return nil
}

func (rh *RepoHandler) GetAllSlots(floorID string) ([]*Models.Slot, error) {
	floor, ok := rh.Floors[floorID]
	if !ok {
		return []*Models.Slot{}, fmt.Errorf("floor not found for id %s: ", floorID)
	}

	return floor.ListSlots, nil
}

func (rh *RepoHandler) AddSlot(slot *Models.Slot) error {
	if slot == nil {
		return errors.New("slot is nil")
	}
	rh.Slots[slot.ID] = slot
	return nil
}

func (rh *RepoHandler) RemoveSlot(slot *Models.Slot) error {
	if slot == nil {
		return errors.New("slot is nil")
	}
	delete(rh.Slots, slot.ID)
	return nil
}

func (rh *RepoHandler) AddTicket(ticket *Models.Ticket) error {
	if ticket == nil {
		return errors.New("ticket is nil")
	}
	rh.Tickets[ticket.ID] = ticket
	return nil
}

func (rh *RepoHandler) AddVehicle(vehicle *Models.Vehicle) error {
	if vehicle == nil {
		return errors.New("vehicle is nil")
	}
	rh.Vehicles[vehicle.ID] = vehicle
	return nil
}
