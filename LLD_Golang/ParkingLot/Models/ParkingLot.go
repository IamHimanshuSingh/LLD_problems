package Models

type ParkingLot struct {
	ID           string   `json:"id"`
	Address      string   `json:"address"`
	ListOfFloors []*Floor `json:"listOfFloors"`
}

func NewParkingLot(address string, ID string) *ParkingLot {
	return &ParkingLot{Address: address, ID: ID, ListOfFloors: make([]*Floor, 0)}
}
