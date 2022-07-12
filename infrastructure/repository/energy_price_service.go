package repository

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"rpi-house-energy-display/domain/contracts"
	"rpi-house-energy-display/domain/model"
	"time"
)

type precioLuzEnergyPoolService struct {
	client *http.Client
}

func (p *precioLuzEnergyPoolService) GetCurrentEnergyCost() (*model.Cost, error) {
	var currentCostResponse currentEnergyCost
	req, error := http.NewRequest("GET", "https://api.preciodelaluz.org/v1/prices/now?zone=PCB", nil)
	if error != nil {
		logrus.Error(error)
		return nil, error
	}

	response, _ := p.client.Do(req)
	if response != nil {
		defer response.Body.Close()
	}

	respBody, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(respBody, &currentCostResponse)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &model.Cost{
		Value: currentCostResponse.Price / 1000000, // convert price to eur/watts
		Date:  time.Time{},
	}, nil

}

func NewEnergyPoolService(client *http.Client) contracts.EnergyPoolService {
	return &precioLuzEnergyPoolService{client: client}
}

type currentEnergyCost struct {
	Date       string  `json:"date"`
	Hour       string  `json:"hour"`
	IsCheap    bool    `json:"is-cheap"`
	IsUnderAvg bool    `json:"is-under-avg"`
	Market     string  `json:"market"`
	Price      float64 `json:"price"`
	Units      string  `json:"units"` //unit is Euro/Mega Watts
}
