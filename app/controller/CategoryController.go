package controller

import (
	"github.com/gin-gonic/gin"
	"goML/app/clients"
	"goML/app/model"
	"log"
	"net/http"
)

func HandlerEmptyGen(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "category ID not provided",
	})
}

func HandlerCategory(c *gin.Context) {
	var err error
	var gen *model.CategoryForGen
	catID := c.Param("catID")

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
