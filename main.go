package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-sdk/sdk"
	"log"
	"net/http"
	"strings"
	"sync"
)

const (
	ClientID     int64  = 0
	ClientSecret string = ""
	Host         string = "https://localhost:8080"
)

var (
	wg         sync.WaitGroup
	chSite     = make(chan *Site)
	chSeller   = make(chan *Seller)
	chCategory = make(chan *Category)
)

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

	r.GET("/show/:itemID", func(c *gin.Context) {
		itemID := c.Param("itemID")
		fmt.Println("item received", itemID)

		if itemID != "" {
			if item, err = getItemByID(client, itemID); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				c.JSON(http.StatusNotFound, nil)
			} else {
				attributes := strings.Split(c.Query("attributes"), ",")
				result = getResultJson(client, item, attributes)
				c.JSON(http.StatusOK, result)
			}
		}
	})

	r.GET("/genealogy/:catID",func(c *gin.Context) {
		catID := c.Param("catID")
		fmt.Println("Category received", catID)
		if catID != "" {
			if gen, err := getGenealogyByID(client, catID); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				c.JSON(http.StatusNotFound, nil)
			} else {
				fmt.Println("genealogy:", gen)
				//TODO armar json de respuesta
				//c.JSON(http.StatusOK, result)
			}
		}

	})

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Cannot run:", error.Error(err))
	}
}



func getResultJson(client *sdk.Client, item *Item, attributes []string) itemInfo {

	var site *Site
	var seller *Seller
	var category *Category
	noAttrs := attributes[0] == ""
	attrSite := paramObtained(attributes, "site")
	attrSeller := paramObtained(attributes, "seller")
	attrCategory := paramObtained(attributes, "category")

	if noAttrs {
		wg.Add(3)
	} else {
		wg.Add(len(attributes))
	}

	// collect attributes
	if noAttrs || attrSite {
		go getSiteByID(client, item.SiteId)
	}
	if noAttrs || attrSeller {
		go getSellerByID(client, item.SellerId)
	}
	if noAttrs || attrCategory {
		go getCategoryByID(client, item.CategoryId)
	}

	//wait for channels
	if noAttrs || attrSite {
		site = <-chSite
	}
	if noAttrs || attrSeller {
		seller = <-chSeller
	}
	if noAttrs || attrCategory {
		category = <-chCategory
	}
	// Wait until everyone finishes.
	wg.Wait()
	return getMergedResults(item, site, seller, category)
}

func paramObtained(attributes []string, s string) bool {
	for i := range attributes {
		if s == strings.ToLower(attributes[i]) {
			fmt.Println("found attr", s)
			return true
		}
	}
	fmt.Println("not found attr", s)
	return false
}

func setupRouter() *gin.Engine {
	// gin.DisableConsoleColor()
	r := gin.Default()
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("",func(c*gin.Context){})
	return r
}
