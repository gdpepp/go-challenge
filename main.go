package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-sdk/sdk"
	"log"
	"net/http"
	"sync"
)

const (
	ClientID int64 = 0
	ClientSecret string = ""
	Host string = "https://localhost:8080"
	)

var wg sync.WaitGroupagreg
var chSite = make(chan *Site)
var chSeller = make(chan *Seller)
var chCategory = make(chan *Category)

func main() {
	r := setupRouter()
	//setup MELI Client
	client, err := sdk.Meli(ClientID, "", ClientSecret, Host)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}

	var item *Item
	var result itemInfo

	r.GET("/items/:itemID", func(c *gin.Context) {
		itemID := c.Param("itemID")
		if item, err = getItemByID(client, itemID); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			c.JSON(http.StatusNotFound, nil)
		} else {
			result = getResultJson(client, item)
			c.JSON(http.StatusOK, result)
		}
	})

	if err := r.Run(":8080"); err!= nil {
		fmt.Println("Cannot run:", error.Error(err))
	}
}

func getResultJson(client *sdk.Client, item *Item) itemInfo {
	wg.Add(3)
	go getSiteByID(client, item.SiteId)
	go getSellerByID(client, item.SellerId)
	go getCategoryByID(client, item.CategoryId)

	// Wait until everyone finishes.
	site := <- chSite
	seller := <- chSeller
	category:= <- chCategory
	wg.Wait()

	return getMergedResults(item, site, seller, category)
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