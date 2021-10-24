package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// const PATH_API string = "http://10.104.0.14:7071/"

var PATH_API string = "http://34.126.177.103/"

func GetApiPath() string {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	appsApi := os.Getenv("PATH_API")
	if appsApi == "" {
		appsApi = "http://34.126.177.103/"
	}

	return appsApi
}
