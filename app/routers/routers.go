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
	return r
}
