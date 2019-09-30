package model

import (
	"encoding/json"
	"time"
)

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
