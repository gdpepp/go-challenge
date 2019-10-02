package routers

import (
	"github.com/gin-gonic/gin"
	"goML/app/controller"
	"goML/app/urls"
)

func GetRouter() *gin.Engine {
	urlPatterns := urls.ReturnURLS()
	r := gin.Default()
	r.GET(urlPatterns.SHOW_PATH, controller.HandlerProduct)
	r.GET(urlPatterns.GENEALOGY, controller.HandlerCategory)
	r.GET(urlPatterns.SHOW_EMPTY, controller.HandlerEmptyProduct)
	r.GET(urlPatterns.GENEALOGY_EMPTY, controller.HandlerEmptyGen)
	return r
}
