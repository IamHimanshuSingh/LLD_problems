package Service

import (
	"ParkingLot/Models"
	"ParkingLot/Repo"
)

type VehicleService interface {
	AddVehicle(vehicle *Models.Vehicle) error
	RemoveVehicle(vehicle *Models.Vehicle) error
}

type VehicleHandler struct {
	Repo *Repo.RepoHandler
}

func NewVehicleHandler(repo *Repo.RepoHandler) *VehicleHandler {
	return &VehicleHandler{Repo: repo}
}

func (vh *VehicleHandler) AddVehicle(vehicle *Models.Vehicle) error {
	return vh.Repo.AddVehicle(vehicle)
}

func (vh *VehicleHandler) RemoveVehicle(vehicle *Models.Vehicle) error {
	return vh.Repo.AddVehicle(vehicle)
}
