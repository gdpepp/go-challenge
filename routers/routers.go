package routers

import (
	"challenge/url"
	"github.com/gin-gonic/gin"
)
func GetRouter() *gin.Engine{
	urlPatterns :=url.ReturnURLS()
	r := gin.Default()
	r.GET(urlPatterns.SHOW_PATH)

	return r
}
