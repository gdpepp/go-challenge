package model

import "time"

type Seller struct {
	SellerId         int32     `json:"seller_id"`
	Id               int32     `json:"id"`
	Nickname         string    `json:"nickname"`
	RegistrationDate time.Time `json:"registration_date"`
}
