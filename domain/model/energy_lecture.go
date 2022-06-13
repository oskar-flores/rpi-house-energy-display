package model

import "time"

type EnergyLecture struct {
	Id           uint
	LectureValue string
	LectureDate  time.Time
}
