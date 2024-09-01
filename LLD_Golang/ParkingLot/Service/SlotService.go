package Service

import (
	"ParkingLot/Models"
	"ParkingLot/Repo"
	"errors"
	"sync"
)

type SlotService interface {
	AddSlot(slot *Models.Slot) error
	RemoveSlot(slot *Models.Slot) error
	UpdateSlotStatus(slotID string, status bool) error
}

type SlotServiceHandler struct {
	vehicleService  *VehicleHandler
	floorSvcHandler *FloorServiceHandler
	Repo            *Repo.RepoHandler
	mu              sync.Mutex
}

func NewSlotServiceHandler(vehicleService *VehicleHandler, repo *Repo.RepoHandler) *SlotServiceHandler {
	return &SlotServiceHandler{vehicleService: vehicleService, Repo: repo}
}

func (sh *SlotServiceHandler) AddSlot(slot *Models.Slot) error {
	sh.Repo.Floors[slot.FloorID].ListSlots = append(sh.Repo.Floors[slot.FloorID].ListSlots, slot)
	return sh.Repo.AddSlot(slot)
}

func (sh *SlotServiceHandler) RemoveSlot(slot *Models.Slot) error {
	return sh.Repo.RemoveSlot(slot)
}

func (sh *SlotServiceHandler) UpdateSlotStatus(slotID string, status bool) error {
	slot, ok := sh.Repo.Slots[slotID]
	if !ok {
		return errors.New("slot Not Found")
	}
	sh.mu.Lock()
	slot.IsOccupied = status
	sh.Repo.Slots[slotID] = slot
	sh.mu.Unlock()
	return nil
}
