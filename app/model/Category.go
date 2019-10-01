package model

type Category struct {
	CategoryId               int64  `json:"category_id"`
	Id                       string `json:"id"`
	Name                     string `json:"name"`
	TotalItemsInThisCategory int    `json:"total_items_in_this_category"`
	Picture                  string `json:"picture"`
}