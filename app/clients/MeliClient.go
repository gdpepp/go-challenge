package clients

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-sdk/sdk"
	"goML/app/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var Client *sdk.Client

var (
	Wg         sync.WaitGroup
	ChSite     = make(chan *model.Site)
	ChSeller   = make(chan *model.Seller)
	ChCategory = make(chan *model.Category)
)

func FetchItem(itemID string) (*model.Item, error) {
	var response *http.Response
	var err error
	var item = new(model.Item)

	if response, err = Client.Get("/items/" + itemID); err != nil {
		fmt.Println("Error:", err.Error())
		return nil, err
	}
	jsonBytes, _ := ioutil.ReadAll(response.Body)

	if err = json.Unmarshal(jsonBytes, item); err != nil {
		return nil, err
	}

	return item, nil
}

func FetchSiteByID(client *sdk.Client, siteID string) {
	defer Wg.Done()
	var response *http.Response
	var err error
	var sit = new(model.Site)
	if response, err = client.Get("/sites/" + siteID); err != nil {
		log.Printf("Error en Get: %s\n", err.Error())
		//return nil, err
	} else {
		jsonBytes, _ := ioutil.ReadAll(response.Body)

		if err = json.Unmarshal(jsonBytes, sit); err != nil {
			log.Printf("Error en Unmarshall site: %s\n", err.Error())
		}
		//return site, nil
		ChSite <- sit
	}
}

func FetchSellerByID(client *sdk.Client, sellerID int32) {
	defer Wg.Done()
	var response *http.Response
	var err error
	var sel = new(model.Seller)

	if response, err = client.Get("/users/" + strconv.FormatInt(int64(sellerID), 10)); err != nil {
		log.Printf("Error en Get: %s\n", err.Error())
		//return nil, err
	} else {
		jsonBytes, _ := ioutil.ReadAll(response.Body)

		if err = json.Unmarshal(jsonBytes, sel); err != nil {
			log.Printf("Error en Unmarshall: %s\n", err.Error())
			//return nil, err
		}

		//return seller, nil
		ChSeller <- sel
	}
}

func FetchCategoryByID(client *sdk.Client, categoryID string) {
	defer Wg.Done()
	var response *http.Response
	var err error
	var cat = new(model.Category)

	if response, err = client.Get("/categories/" + categoryID); err != nil {
		log.Printf("Error en Get: %s\n", err.Error())
	} else {
		jsonBytes, _ := ioutil.ReadAll(response.Body)

		if err = json.Unmarshal(jsonBytes, cat); err != nil {
			log.Printf("Error en Unmarshall: %s\n", err.Error())
		}
		ChCategory <- cat
	}
}