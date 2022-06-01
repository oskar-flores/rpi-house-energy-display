package contracts

import "rpi-house-energy-display/domain/model"

type EnergyPoolService interface {
	GetCurrentEnergyCost() (model.Cost, error)
}
