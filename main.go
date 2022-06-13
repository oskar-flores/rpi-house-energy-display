package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"rpi-house-energy-display/infrastructure/config"
	"rpi-house-energy-display/infrastructure/repository"
)

func main() {
	configuration := config.NewConfig(os.Getenv("user"), os.Getenv("pass"))

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}

	client := http.Client{Jar: jar}
	repo := repository.NewIdeMeterRepository(&client, *configuration)

	print(repo.GetCurrentLecture())

}
