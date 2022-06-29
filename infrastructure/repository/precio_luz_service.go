package repository

import (
	"net/http"
	"rpi-house-energy-display/domain/contracts"
	"rpi-house-energy-display/domain/model"
)

type precioLuzEnergyPoolService struct {
	client *http.Client
}

func (p *precioLuzEnergyPoolService) GetCurrentEnergyCost() (model.Cost, error) {
	//TODO implement me
	panic("implement me")
}

func NewEnergyPoolService(client *http.Client) contracts.EnergyPoolService {
	return &precioLuzEnergyPoolService{client: client}
}
