package main

import (
	"fmt"
	cookiejar "github.com/orirawlings/persistent-cookiejar"
	"log"
	"net/http"
	"os"
	"rpi-house-energy-display/domain/model"
	"rpi-house-energy-display/infrastructure/config"
	"rpi-house-energy-display/infrastructure/repository"
)

func main() {
	configuration := config.NewConfig(os.Getenv("IDE_USER"), os.Getenv("IDE_PASS"))
	lecturesChannel := make(chan *model.EnergyLecture, 1)

	jar, err := cookiejar.New(&cookiejar.Options{})
	defer jar.Save()
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}

	client := http.Client{Jar: jar}
	repo := repository.NewIdeMeterRepository(&client, *configuration)
	go repo.GetCurrentLecture(lecturesChannel)

	currentLectureValue := <-lecturesChannel

	fmt.Println(currentLectureValue)

}
