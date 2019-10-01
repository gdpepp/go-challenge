package model

type GenealogyResponse struct {
	Category           	Category
	Roots              	[]Category
	Children_categories []Category
}
