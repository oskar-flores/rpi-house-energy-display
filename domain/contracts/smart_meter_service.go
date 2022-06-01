package contracts

import model "rpi-house-energy-display/domain/model"

type SmartMeterService interface {
	GetCurrentLecture() (model.EnergyLecture, error)
}
