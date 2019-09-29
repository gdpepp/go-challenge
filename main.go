package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-sdk/sdk"
	"io/ioutil"
	"log"
	"net/http"
)

const ClientID int64 = 0
const ClientSecret string = ""

type itemInfo struct {
	Id                string       `json:"id"`
	CategoryId        string       `json:"category_id"`
	Title             string       `json:"title"`
	SellerId          int64        `json:"seller_id"`
	OfficialStoreId   *int64       `json:"official_store_id"`
	Price             *json.Number `json:"price"`
	BasePrice         *json.Number `json:"base_price"`
	CurrencyId        *string      `json:"currency_id"`
	InitialQuantity   int32        `json:"initial_quantity"`
	AvailableQuantity int32        `json:"available_quantity"`
	AcceptsMercadoPago bool         `json:"accepts_mercadopago"`
}

func main() {

	r := setupRouter()

	//setup MELI Client
	client, err := sdk.Meli(ClientID, "", ClientSecret, "https://localhost:8080")
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}

	r.GET("/items/:itemID", func(c *gin.Context) {
		itemID := c.Param("itemID")
		var itemInfo *itemInfo
		if itemInfo, err = getItemByID(client, itemID); err == nil {
			//respond with item json
			c.JSON(http.StatusOK, gin.H{"message" : itemInfo})
		} else {
			c.JSON(http.StatusNotFound, nil)
		}
	})

	if err := r.Run(":8080"); err!= nil {
		fmt.Println("Cannot run:", error.Error(err))
	}
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

func getItemByID(client *sdk.Client, itemID string) (*itemInfo, error) {
	var response *http.Response
	var err error
	var item = new(itemInfo)

	if response, err = client.Get("/items/" + itemID); err != nil {
		log.Printf("Error: %s\n", err.Error())
		return nil, err
	}
	jsonBytes, _ := ioutil.ReadAll(response.Body)

	fmt.Printf("Item Info :%s\n", jsonBytes)
	if err = json.Unmarshal(jsonBytes, item); err != nil {
		return nil, err
	}

	return item, nil
}

