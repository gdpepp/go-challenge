package main

import (
	"github.com/mercadolibre/golang-sdk/sdk"
	"goML/app/clients"
	"goML/app/routers"
	"log"
)

const (
	ClientID     int64  = 0
	ClientSecret string = ""
	Host         string = "https://localhost:8080"
)


func main() {
	r := routers.GetRouter()
	client, err := sdk.Meli(ClientID, "", ClientSecret, Host)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}

	clients.Client = client
	if err := r.Run(":8080"); err != nil {
		log.Println("Cannot run:", error.Error(err))
	}
}
