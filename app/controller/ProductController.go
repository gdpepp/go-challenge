package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goML/app/clients"
	"goML/app/model"
	"net/http"
	"strings"
)

func HandlerProduct(c *gin.Context) {
	itemID := c.Param("itemID")

	var item *model.Item
	var err error

	if item, err = clients.FetchItem(itemID); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		c.JSON(http.StatusNotFound, nil)
	} else {
		attributes := strings.Split(c.Query("attributes"), ",")
		result := getResultJson(item, attributes)
		c.JSON(http.StatusOK, result)
	}

}

func getResultJson(item *model.Item, attributes []string) model.ItemResponse {

	var site *model.Site
	var seller *model.Seller
	var category *model.Category
	noAttrs := attributes[0] == ""
	attrSite := paramObtained(attributes, "site")
	attrSeller := paramObtained(attributes, "seller")
	attrCategory := paramObtained(attributes, "category")

	if noAttrs {
		clients.Wg.Add(3)
	} else {
		clients.Wg.Add(len(attributes))
	}

	// collect attributes
	if noAttrs || attrSite {
		go clients.FetchSiteByID(item.SiteId)
	}
	if noAttrs || attrSeller {
		go clients.FetchSellerByID(item.SellerId)
	}
	if noAttrs || attrCategory {
		go clients.FetchCategoryByID(item.CategoryId)
	}

	//wait for channels
	if noAttrs || attrSite {
		site = <-clients.ChSite
	}
	if noAttrs || attrSeller {
		seller = <-clients.ChSeller
	}
	if noAttrs || attrCategory {
		category = <-clients.ChCategory
	}
	// Wait until everyone finishes.
	clients.Wg.Wait()
	return GetMergedResults(item, site, seller, category)
}

func paramObtained(attributes []string, s string) bool {
	for i := range attributes {
		if s == strings.ToLower(attributes[i]) {
			return true
		}
	}
	return false
}

func GetMergedResults(item *model.Item, site *model.Site, seller *model.Seller, category *model.Category) model.ItemResponse {
	var result model.ItemResponse

	if site != nil {
		result.Site = site
	}
	if category != nil {
		result.Category = category
	}
	if seller != nil {
		result.Seller = seller
	}
	result.Id = item.Id
	result.Title = item.Title
	result.Subtitle = item.Subtitle
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
