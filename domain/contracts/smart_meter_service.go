package contracts

import (
	model "rpi-house-energy-display/domain/model"
)

type SmartMeterService interface {
	/*
			Reads the data from the smart meter using a post call
			It takes some time, so we use a channel to make it asynchronous.
		    Current lecture is returned in watts
	*/
	GetCurrentLecture(rc chan *model.EnergyMeasurement) error
}
