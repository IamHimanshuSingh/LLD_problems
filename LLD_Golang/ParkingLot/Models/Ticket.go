package Models

type Ticket struct {
	ID           string `json:"id"`
	SlotID       string `json:"slotId"`
	FloorID      string `json:"floorId"`
	ParkingLotID string `json:"parkingLotId"`
	VehicleID    string `json:"vehicleId"`
	Price        int    `json:"price"`
	IssuedAt     string `json:"issuedAt"`
	ExpireAt     string `json:"expireAt"`
}

// for structs with lots of constructor parameters, 'builder pattern' can be helpful
// In java we get that as part of lombok, but in golang we don't have anything like lombok
// to support such behavior

// Golang has Functional Options Pattern as an alternative:
// Read: https://blog.matthiasbruns.com/golang-options-vs-builder-pattern

func NewTicket(ID string, slotID string, floorID string, parkingLotId string, vehicleID string, price int, issuedAt string, expireAt string) *Ticket {
	return &Ticket{ID: ID, SlotID: slotID, FloorID: floorID, ParkingLotID: parkingLotId, VehicleID: vehicleID, Price: price, IssuedAt: issuedAt, ExpireAt: expireAt}
}
