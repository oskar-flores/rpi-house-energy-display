package model

import "time"

type EnergyMeasurement struct {
	Id           uint
	LectureValue float64
	LectureDate  time.Time
}
