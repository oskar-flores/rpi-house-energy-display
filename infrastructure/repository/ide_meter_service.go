package repository

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"rpi-house-energy-display/domain/contracts"
	"rpi-house-energy-display/domain/model"
)

type ideMeterRepository struct {
	client *http.Client
}

func NewIdeMeterRepository(client *http.Client) contracts.SmartMeterService {
	return &ideMeterRepository{client: client}
}

func (imr *ideMeterRepository) GetCurrentLecture() (model.EnergyLecture, error) {

	panic("implement me")
}

func (imr *ideMeterRepository) login(user string, pass string) {
	var payload = fmt.Sprintf("[%s, %s]", user, pass)
	var jsonStr = []byte(payload)
	request, err := http.NewRequest("POST", "https://www.i-de.es/consumidores/rest/loginNew/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	postResponse, err := imr.client.Do(request)
	if err != nil {
		return
	}
	logrus.Debug(postResponse)
}
