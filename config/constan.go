package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// const PATH_API string = "http://10.104.0.14:7071/"

var PATH_API string = "http://128.199.241.112:7071/"

func GetApiPath() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	appsApi := os.Getenv("PATH_API")
	if appsApi == "" {
		appsApi = "http://128.199.241.112:7071/"
	}

	return appsApi
}

func GetDreamApiPath() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	deramApi := os.Getenv("PATH_PANEL_API")
	if deramApi == "" {
		deramApi = "http://128.199.241.112:7071/"
	}

	return deramApi
}
