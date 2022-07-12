package contracts

import "rpi-house-energy-display/domain/model"

type EnergyPoolService interface {
	/*
		Returns the current price of the energy in Eur/watts
	*/
	GetCurrentEnergyCost() (*model.Cost, error)
}
