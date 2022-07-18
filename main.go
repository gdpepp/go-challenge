package main

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-sdk/sdk"
	"go-challenge/app/clients"
	"go-challenge/app/routers"
	"log"
	"os"
)

type conf struct {
	ClientID     int64
	ClientSecret string
	Host         string
}

func main() {

	conf := getConfParameters()
	r := routers.GetRouter()

	client, err := sdk.Meli(conf.ClientID, "", conf.ClientSecret, conf.Host)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}

	clients.Client = client
	if err := r.Run(":8080"); err != nil {
		log.Println("Cannot run:", error.Error(err))
	}
}

func getConfParameters() conf {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := conf{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
