package Service

import (
	"ParkingLot/Models"
	"ParkingLot/Repo"
	"errors"
	"sync"
)

type ParkingLotService interface {
	AddParkingLot(parkingLot *Models.ParkingLot) error
	RemoveParkingLot(parkingLot *Models.ParkingLot) error
	GetAllParkingLots() error
	ParkVehicle(vehicle *Models.Vehicle, ParkingLotID string) (*Models.Ticket, error)
	UnParkVehicle(ticket *Models.Ticket) error
}

// TODO: dependency injection using Wiring
type ParkingLotServiceHandler struct {
	Repo          *Repo.RepoHandler
	ticketHandler *TicketHandler
	slotHandler   *SlotServiceHandler
	mu            sync.Mutex
}

func (plsh *ParkingLotServiceHandler) GetAllParkingLots() error {
	return plsh.Repo.GetAllParkingLots()
}

func NewParkingLotServiceHandler(repo *Repo.RepoHandler, th *TicketHandler, sh *SlotServiceHandler) *ParkingLotServiceHandler {
	return &ParkingLotServiceHandler{
		Repo:          repo,
		ticketHandler: th,
		slotHandler:   sh,
	}
}

func (plsh *ParkingLotServiceHandler) AddParkingLot(parkingLot *Models.ParkingLot) error {
	return plsh.Repo.AddParkingLot(parkingLot)
}

func (plsh *ParkingLotServiceHandler) RemoveParkingLot(parkingLot *Models.ParkingLot) error {
	return plsh.Repo.RemoveParkingLot(parkingLot)
}

func (plsh *ParkingLotServiceHandler) ParkVehicle(vehicle *Models.Vehicle, parkingLotID string) (*Models.Ticket, error) {
	// TODO: Find available parking slot for given vehicle type using strategy
	parkingLot, ok := plsh.Repo.ParkingLots[parkingLotID]
	if !ok {
		return nil, errors.New("parking lot not exist")
	}

	var availableSlot *Models.Slot
	for _, floor := range parkingLot.ListOfFloors {
		for _, slot := range floor.ListSlots {
			if !slot.IsOccupied && slot.Type == vehicle.Type {
				availableSlot = slot
				break
			}
		}
	}

	if availableSlot == nil {
		return nil, errors.New("parking slot not available for given vehicle")
	}

	// Generate Ticket
	ticket, err := plsh.ticketHandler.GenerateTicket(vehicle, parkingLot, availableSlot)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (plsh *ParkingLotServiceHandler) UnParkVehicle(ticket *Models.Ticket) error {
	// TODO: Update the slot, check for concurrency
	plsh.mu.Lock()
	err := plsh.slotHandler.UpdateSlotStatus(ticket.SlotID, false)
	if err != nil {
		return err
	}
	plsh.mu.Unlock()
	return nil
}
