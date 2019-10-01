package model

import (
	"encoding/json"
	"time"
)

type ItemResponse struct {
	Id                string
	Site              *Site
	Title             string
	Subtitle          string
	Seller            *Seller
	Category          *Category
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