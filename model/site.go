package model

type Site struct {
	SiteId             string `json:"site_id"`
	Id                 string `json:"id"`
	CountryId          string `json:"country_id"`
	SaleFeesMode       string `json:"sale_fees_mode"`
	MercadopagoVersion int    `json:"mercadopago_version"`
}