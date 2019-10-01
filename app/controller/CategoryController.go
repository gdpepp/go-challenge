package controller

import (
	"github.com/gin-gonic/gin"
	"goML/app/clients"
	"goML/app/model"
	"log"
	"net/http"
)

func HandlerCategory(c *gin.Context) {
	catID := c.Param("catID")

	var gen *model.CategoryForGen
	var err error

	if gen, err = clients.FetchGenealogy(catID); err != nil {
		log.Printf("Error: %s\n", err.Error())
		c.JSON(http.StatusNotFound, nil)
	} else {
		result := GetMergedCategoryResults(gen)
		c.JSON(http.StatusOK, result)
	}

}

func GetMergedCategoryResults(gen *model.CategoryForGen) model.GenealogyResponse {
	var result model.GenealogyResponse

	result.Category.Id = gen.Id
	result.Category.CategoryId = gen.CategoryId
	result.Category.Name = gen.Name
	result.Category.Picture = gen.Picture
	result.Category.TotalItemsInThisCategory = gen.TotalItemsInThisCategory

	if gen.Children != nil {
		result.Children_categories = gen.Children
	}
	if gen.Parents != nil {
		result.Roots = gen.Parents
	}

	return result
}