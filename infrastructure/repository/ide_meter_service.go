package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"rpi-house-energy-display/domain/contracts"
	"rpi-house-energy-display/domain/model"
	"rpi-house-energy-display/infrastructure/config"
	"time"
)

type ideMeterRepository struct {
	client *http.Client
	config config.Config
}

func NewIdeMeterRepository(client *http.Client, config config.Config) contracts.SmartMeterService {
	return &ideMeterRepository{client: client, config: config}
}

func (imr *ideMeterRepository) GetCurrentLecture() (model.EnergyLecture, error) {

	if imr.validateConnection() {
		var read readResponse

		request, err := http.NewRequest("GET", "https://www.i-de.es/consumidores/rest/escenarioNew/obtenerMedicionOnline/24", nil)

		if err != nil {
			return model.EnergyLecture{}, err
		}
		request.Header.Set("AppVersion", "v2")
		request.Header.Set("Accept", "application/json, text/plain, */*")
		request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Safari/537.36")
		request.Header.Set("dispositivo", "desktop")
		getRespose, err := imr.client.Do(request)
		if err != nil {
			logrus.Error(getRespose)
		}

		if getRespose != nil {
			defer getRespose.Body.Close()
		}

		respBody, _ := ioutil.ReadAll(getRespose.Body)
		err = json.Unmarshal(respBody, &read)
		if err != nil {
			return model.EnergyLecture{}, err
		}

		return model.EnergyLecture{
			Id:           0,
			LectureValue: read.ValLecturaContador,
			LectureDate:  time.Now(),
		}, nil

	}

	return model.EnergyLecture{}, errors.New("Unable to validate connection to the smart meter ")
}
func (imr *ideMeterRepository) login(user string, pass string) {
	var payload = fmt.Sprintf("[\"%s\",\"%s\",\"null\", \"Mac OS X 10_15_7\",\"PC\",\"Chrome 102.0.5005.115\",\"0\",\"\",\"n\"]", user, pass)
	var jsonStr = []byte(payload)
	request, _ := http.NewRequest("POST", "https://www.i-de.es/consumidores/rest/loginNew/login", bytes.NewBuffer(jsonStr))

	request.Header.Set("AppVersion", "v2")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Safari/537.36")
	request.Header.Set("dispositivo", "desktop")
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	postResponse, err := imr.client.Do(request)
	if postResponse != nil {
		defer postResponse.Body.Close()
		printResponse(postResponse)
	}

	if err != nil {
		return
	}
	logrus.Debug(postResponse)
}

func (imr *ideMeterRepository) validateConnection() bool {
	imr.login(imr.config.User(), imr.config.Password())
	//first validate conection
	request, err := http.NewRequest("GET", "https://www.i-de.es/consumidores/rest/escenarioNew/validarComunicacionContador/", nil)

	if err != nil {
		logrus.Error(err)
		return false
	}
	request.Header.Set("AppVersion", "v2")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Safari/537.36")
	request.Header.Set("dispositivo", "desktop")

	getRespose, err := imr.client.Do(request)
	if getRespose != nil {
		defer getRespose.Body.Close()
		printResponse(getRespose)
	}

	if err != nil {
		logrus.Error(getRespose)
	}
	return true

}

type readResponse struct {
	ValMagnitud        string
	ValInterruptor     string
	ValEstado          string
	ValLecturaContador string
	CodSolicitudTGT    string
}

func printResponse(response *http.Response) {
	body, _ := ioutil.ReadAll(response.Body)
	dst := &bytes.Buffer{}
	_ = json.Indent(dst, body, "", "")
	fmt.Println(dst.String())
}
