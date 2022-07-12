package main

import (
	"fmt"
	cookiejar "github.com/orirawlings/persistent-cookiejar"
	"log"
	"net/http"
	"os"
	"rpi-house-energy-display/apps"
	outputModels "rpi-house-energy-display/apps/model"
	"rpi-house-energy-display/domain/model"
	"rpi-house-energy-display/infrastructure/config"
	"rpi-house-energy-display/infrastructure/repository"
)

func main() {
	configuration := config.NewConfig(os.Getenv("IDE_USER"), os.Getenv("IDE_PASS"))
	lecturesChannel := make(chan *model.EnergyMeasurement, 1)

	jar, err := cookiejar.New(&cookiejar.Options{})
	defer jar.Save()
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}

	client := http.Client{Jar: jar}
	meterService := repository.NewIdeMeterService(&client, *configuration)
	priceService := repository.NewEnergyPoolService(&client)

	go meterService.GetCurrentLecture(lecturesChannel)
	price, _ := priceService.GetCurrentEnergyCost()

	display := apps.Newwavesahre213Display()
	defer display.Close()
	defer display.Epd.TurnDisplayOff()

	currentLectureValue := <-lecturesChannel
	displayModel, err := outputModels.NewDisplayModel(price.Value, currentLectureValue.LectureValue)
	if err != nil {
		return
	}
	display.Draw(*displayModel)

	fmt.Println(currentLectureValue)

}
