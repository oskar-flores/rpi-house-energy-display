package model

import "errors"

type DisplayModel struct {
	CurrentPrice     float64
	CurrentLecture   float64
	CurrentCostInPVC float64
}

func NewDisplayModel(price, lecture float64) (*DisplayModel, error) {
	if price == 0 || lecture == 0 {
		return nil, errors.New("price or lecture cannot be 0")
	}

	return &DisplayModel{
		CurrentPrice:     price,
		CurrentLecture:   lecture,
		CurrentCostInPVC: price * lecture,
	}, nil

}
