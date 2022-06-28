package main

import (
	"fmt"
	cookiejar "github.com/orirawlings/persistent-cookiejar"
	"github.com/oskar-flores/edp_2.13_V3"
	"log"
	"net/http"
	"os"
	"rpi-house-energy-display/apps"
	"rpi-house-energy-display/domain/model"
	"rpi-house-energy-display/infrastructure/config"
	"rpi-house-energy-display/infrastructure/repository"
	"time"
)

func main() {
	test()
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
	display := apps.Newwavesahre213Display()
	defer display.Close()

	currentLectureValue := <-lecturesChannel
	display.Draw(currentLectureValue)
	display.Epd.Display(getData())

	fmt.Println(currentLectureValue)

}

func test() {
	e := epaper.CreateEpd()
	defer e.Close()
	defer e.Clear()
	e.Init()
	e.Clear()

	fmt.Printf("Display\n")
	e.Display(getData())
	fmt.Printf("sleeping\n")
	time.Sleep(5 * time.Second)
}

// getData return a image from the gotchi project, more info at https://github.com/GaelicThunder/Chao-Pi-Adventure
func getData() []byte {
	return []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 231, 255, 206, 115, 255, 63, 255, 143, 255, 255, 255, 255, 255, 255, 255, 255, 231, 255, 206, 115, 255, 63, 255, 143, 255, 255, 255, 255, 255, 255, 255, 255, 231, 254, 49, 131, 255, 63, 254, 127, 255, 255, 255, 255, 255, 255, 255, 255, 231, 254, 49, 131, 255, 63, 254, 127, 255, 255, 255, 255, 255, 255, 255, 255, 231, 254, 49, 131, 255, 63, 254, 127, 255, 255, 255, 255, 255, 255, 255, 255, 248, 63, 206, 12, 255, 63, 254, 127, 255, 255, 255, 255, 255, 255, 255, 255, 248, 63, 206, 12, 255, 63, 254, 127, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 0, 112, 0, 1, 241, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 0, 112, 0, 1, 241, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 0, 112, 0, 1, 241, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 249, 255, 140, 31, 63, 241, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 249, 255, 140, 31, 63, 241, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 254, 14, 112, 231, 63, 207, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 254, 14, 112, 231, 63, 207, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 254, 14, 112, 231, 63, 207, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 240, 3, 0, 6, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 240, 3, 0, 6, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 143, 231, 62, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 143, 231, 62, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 143, 231, 62, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 240, 24, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 240, 24, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 231, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 231, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 231, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 7, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 7, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 254, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 254, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 254, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 1, 240, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 1, 240, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 48, 3, 131, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 48, 3, 131, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 48, 3, 131, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 0, 0, 0, 124, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 0, 0, 0, 124, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 63, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 63, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 63, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 7, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 7, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 7, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 0, 3, 131, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 0, 3, 131, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 0, 3, 131, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 0, 3, 131, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 7, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 7, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 7, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 63, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 63, 240, 0, 0, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 0, 124, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 0, 124, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 0, 124, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 48, 3, 131, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 48, 3, 131, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 0, 0, 0, 48, 3, 131, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 1, 240, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 1, 240, 3, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 62, 1, 254, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 62,
		1, 254, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 62, 1, 254, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 63, 192, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 63, 192, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 31, 192, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 31, 192, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 31, 192, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 252, 7, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 252, 7, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 31, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 31, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 31, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 7, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 7, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 7, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 231, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 231, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 240, 24, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 240, 24, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 240, 24, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 143, 231, 62, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 143, 231, 62, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 240, 3, 0, 6, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 240, 3, 0, 6, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 240, 3, 0, 6, 63, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 254, 14, 112, 231, 63, 207, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 254, 14, 112, 231, 63, 207, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 249, 255, 140, 31, 63, 241, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 249, 255, 140, 31, 63, 241, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 249, 255, 140, 31, 63, 241, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 0, 112, 0, 1, 241, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 192, 0, 112, 0, 1, 241, 255, 255, 255, 255, 255, 255, 255, 255, 255, 248, 63, 206, 12, 255, 63, 254, 127, 255, 255, 255, 255, 255, 255, 255, 255, 248, 63, 206, 12, 255, 63, 254, 127, 255, 255, 255, 255, 255, 255, 255, 255, 248, 63, 206, 12, 255, 63, 254, 127, 255, 255, 255, 255, 255, 255, 255, 255, 231, 254, 49, 131, 255, 63, 254, 127, 255, 255, 255, 255, 255, 255, 255, 255, 231, 254, 49, 131, 255, 63, 254, 127, 255, 255, 255, 255, 255, 255, 255, 255, 231, 255, 206, 115, 255, 63, 255, 143, 255, 255, 255, 255, 255, 255, 255, 255, 231, 255, 206, 115, 255, 63, 255, 143, 255, 255, 255, 255, 255, 255, 255, 255, 231, 255, 206, 115, 255, 63, 255, 143, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}
}
