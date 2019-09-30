package main

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-sdk/sdk"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func getItemByID(client *sdk.Client, itemID string) (*Item, error) {
	var response *http.Response
	var err error
	var item = new(Item)

	if response, err = client.Get("/items/" + itemID); err != nil {
		fmt.Println("Error:", err.Error())
		return nil, err
	}
	jsonBytes, _ := ioutil.ReadAll(response.Body)

	if err = json.Unmarshal(jsonBytes, item); err != nil {
		return nil, err
	}

	return item, nil
}

func getSiteByID(client *sdk.Client, siteID string) {
	defer wg.Done()
	var response *http.Response
	var err error
	var sit = new(Site)
	if response, err = client.Get("/sites/" + siteID); err != nil {
		log.Printf("Error en Get: %s\n", err.Error())
		//return nil, err
	}
	jsonBytes, _ := ioutil.ReadAll(response.Body)

	if err = json.Unmarshal(jsonBytes, sit); err != nil {
		log.Printf("Error en Unmarshall site: %s\n", err.Error())
	}
	//return site, nil
	chSite <- sit
}

func getSellerByID(client *sdk.Client, sellerID int32)  {
	defer wg.Done()
	var response *http.Response
	var err error
	var sel = new(Seller)

	response, err = client.Get("/users/" + strconv.FormatInt(int64(sellerID), 10))

	if err != nil {
		log.Printf("Error en Get: %s\n", err.Error())
		//return nil, err
	}
	jsonBytes, _ := ioutil.ReadAll(response.Body)

	if err = json.Unmarshal(jsonBytes, sel); err != nil {
		log.Printf("Error en Unmarshall: %s\n", err.Error())
		//return nil, err
	}

	//return seller, nil
	chSeller <- sel
}

func getCategoryByID(client *sdk.Client, categoryID string)  {
	defer wg.Done()
	var response *http.Response
	var err error
	var cat = new(Category)

	if response, err = client.Get("/categories/" + categoryID); err != nil {
		log.Printf("Error en Get: %s\n", err.Error())
		//return nil, err
	}
	jsonBytes, _ := ioutil.ReadAll(response.Body)

	if err = json.Unmarshal(jsonBytes, cat); err != nil {
		log.Printf("Error en Unmarshall: %s\n", err.Error())
		//return nil, err
	}

	chCategory <- cat
	//return category, nil
}

func getMergedResults(item *Item, site *Site, seller *Seller, category *Category) itemInfo {
	var result itemInfo
	result.Id = item.Id
	result.Site = *site
	result.Title = item.Title
	result.Subtitle = item.Subtitle
	result.Seller = *seller
	result.Category = *category
	result.Price = item.Price
	result.BasePrice = item.BasePrice
	result.OriginalPrice = item.OriginalPrice
	result.CurrencyId = item.CurrencyId
	result.InitialQuantity = item.InitialQuantity
	result.AvailableQuantity = item.AvailableQuantity
	result.SoldQuantity = item.SoldQuantity
	result.DateCreated = item.DateCreated
	result.LastUpdated = item.LastUpdated

	return result
}
