package Models

type Floor struct {
	ID           string
	ParkingLotID string
	ListSlots    []*Slot
}

func NewFloor(id string, parkingLotID string, listSlots []*Slot) *Floor {
	return &Floor{ID: id, ParkingLotID: parkingLotID, ListSlots: listSlots}
}
