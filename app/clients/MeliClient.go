package clients

import (
	"encoding/json"
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
	var jsonBytes []byte

	if response, err = Client.Get("/items/" + itemID); err != nil {
		log.Printf("Error: %s\n", err.Error())
		return nil, err
	}
	if jsonBytes, err = ioutil.ReadAll(response.Body); err != nil {
		log.Printf("Error en ReadAll: %s\n", err.Error())
		return nil, err
	}
	if err = json.Unmarshal(jsonBytes, item); err != nil {
		log.Printf("Error en Unmarshal item: %s\n", err.Error())
		return nil, err
	}
	return item, nil
}

func FetchSiteByID(siteID string) {
	defer Wg.Done()
	var response *http.Response
	var err error
	var sit = new(model.Site)
	var jsonBytes []byte

	if response, err = Client.Get("/sites/" + siteID); err != nil {
		log.Printf("Error en Get: %s\n", err.Error())
	} else {
		if jsonBytes, err = ioutil.ReadAll(response.Body); err != nil {
			log.Printf("Error en ReadAll: %s\n", err.Error())
		}
		if err = json.Unmarshal(jsonBytes, sit); err != nil {
			log.Printf("Error en Unmarshal site: %s\n", err.Error())
		}
		ChSite <- sit
	}
}

func FetchSellerByID(sellerID int32) {
	defer Wg.Done()
	var response *http.Response
	var err error
	var sel = new(model.Seller)
	var jsonBytes []byte

	if response, err = Client.Get("/users/" + strconv.FormatInt(int64(sellerID), 10)); err != nil {
		log.Printf("Error en Get: %s\n", err.Error())
	} else {
		if jsonBytes, err = ioutil.ReadAll(response.Body); err != nil {
			log.Printf("Error en ReadAll: %s\n", err.Error())
		}
		if err = json.Unmarshal(jsonBytes, sel); err != nil {
			log.Printf("Error en Unmarshall: %s\n", err.Error())
		}
		ChSeller <- sel
	}
}

func FetchCategoryByID(categoryID string) {
	defer Wg.Done()
	var response *http.Response
	var err error
	var cat = new(model.Category)
	var jsonBytes []byte

	if response, err = Client.Get("/categories/" + categoryID); err != nil {
		log.Printf("Error en Get: %s\n", err.Error())
	} else {
		if jsonBytes, err = ioutil.ReadAll(response.Body); err != nil {
			log.Printf("Error en ReadAll: %s\n", err.Error())
		}
		if err = json.Unmarshal(jsonBytes, cat); err != nil {
			log.Printf("Error en Unmarshall: %s\n", err.Error())
		}
		ChCategory <- cat
	}
}

func FetchGenealogy(categoryID string) (*model.CategoryForGen, error) {
	var response *http.Response
	var err error
	var catg = new(model.CategoryForGen)
	var jsonBytes []byte

	if response, err = Client.Get("/categories/" + categoryID); err != nil {
		log.Printf("Error en Get: %s\n", err.Error())
		return nil, err
	} else {
		if jsonBytes, err = ioutil.ReadAll(response.Body); err != nil {
			log.Printf("Error en ReadAll: %s\n", err.Error())
			return nil, err
		}

		if err = json.Unmarshal(jsonBytes, catg); err != nil {
			log.Printf("Error en Unmarshal: %s\n", err.Error())
			return nil, err
		}
	}
	return catg, nil
}