package Service

import (
	"ParkingLot/Models"
	"ParkingLot/Repo"
	"fmt"
	"sync"
)

type TicketService interface {
	GenerateTicket(vehicle *Models.Vehicle, pLot *Models.ParkingLot, slot *Models.Slot) (*Models.Ticket, error)
}

type TicketHandler struct {
	Repo        *Repo.RepoHandler
	slotHandler *SlotServiceHandler
	mu          sync.Mutex
}

func NewTicketHandler(repo *Repo.RepoHandler, slotHandler *SlotServiceHandler) *TicketHandler {
	return &TicketHandler{Repo: repo, slotHandler: slotHandler}
}

func (th *TicketHandler) GenerateTicket(vehicle *Models.Vehicle, pLot *Models.ParkingLot, slot *Models.Slot) (*Models.Ticket, error) {
	ticket := Models.NewTicket(fmt.Sprintf("Ticket-%s-%s", pLot.ID, slot.ID), slot.ID, slot.FloorID, pLot.ID, vehicle.ID, 100, "2024-02-12:10-00am", "2024-02-12:14-00pm")

	//TODO: concurrency handling required here
	th.mu.Lock()
	err := th.slotHandler.UpdateSlotStatus(slot.ID, true)
	if err != nil {
		return nil, err
	}
	th.mu.Unlock()

	return ticket, nil
}
