package main

import (
	"encoding/json"
	"time"
)

type Site struct {
	SiteId             string `json:"site_id"`
	Id                 string `json:"id"`
	CountryId          string `json:"country_id"`
	SaleFeesMode       string `json:"sale_fees_mode"`
	MercadopagoVersion int    `json:"mercadopago_version"`
}

type Seller struct {
	SellerId         int32     `json:"seller_id"`
	Id               int32     `json:"id"`
	Nickname         string    `json:"nickname"`
	RegistrationDate time.Time `json:"registration_date"`
}
type Category struct {
	CategoryId               int64  `json:"category_id"`
	Id                       string `json:"id"`
	Name                     string `json:"name"`
	TotalItemsInThisCategory int    `json:"total_items_in_this_category"`
	Picture                  string `json:"picture"`
}

type Item struct {
	Id                string       `json:"id"`
	SiteId            string       `json:"site_id"`
	CategoryId        string       `json:"category_id"`
	Title             string       `json:"title"`
	Subtitle          string       `json:"subtitle"`
	SellerId          int32        `json:"seller_id"`
	OfficialStoreId   *int64       `json:"official_store_id"`
	Price             *json.Number `json:"price"`
	BasePrice         *json.Number `json:"base_price"`
	OriginalPrice     *json.Number `json:"original_price"`
	CurrencyId        *string      `json:"currency_id"`
	InitialQuantity   int32        `json:"initial_quantity"`
	AvailableQuantity int32        `json:"available_quantity"`
	SoldQuantity      int32        `json:"sold_quantity"`
	DateCreated       time.Time    `json:"date_created"`
	LastUpdated       time.Time    `json:"last_updated"`
}

//formato de JSON devuelto
type itemInfo struct {
	Id                string
<<<<<<< HEAD
	Site              *Site
	Title             string
	Subtitle          string
	Seller            *Seller
	Category          *Category
=======
	Site              Site
	Title             string
	Subtitle          string
	Seller            Seller
	Category          Category
>>>>>>> fd89830bf86910a0520edc9fca53bc6f899b298b
	Price             *json.Number
	BasePrice         *json.Number
	OriginalPrice     *json.Number
	CurrencyId        *string
	InitialQuantity   int32
	AvailableQuantity int32
	SoldQuantity      int32
	DateCreated       time.Time
	LastUpdated       time.Time
}
