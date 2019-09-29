package main

import (
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-sdk/sdk"
	"io/ioutil"
	"log"
	"net/http"
)

const ClientID int64 = 0
const ClientSecret string = ""

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
		var itemInfo string
		if itemInfo, err = getItemByID(client, itemID); err == nil {
			//respond with item json
			c.JSON(http.StatusOK, itemInfo)
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

func getItemByID(client *sdk.Client, itemID string) (string, error) {
	var response *http.Response
	var err error

	if response, err = client.Get("/items/" + itemID); err != nil {
		log.Printf("Error: %s\n", err.Error())
		return "", err
	}
	jsonBytes, _ := ioutil.ReadAll(response.Body)

	fmt.Printf("Item Info :%s\n", jsonBytes)
	//TODO unmarshal
	//json.Unmarshal(jsonBytes, ? )

	return string(jsonBytes), nil
}

