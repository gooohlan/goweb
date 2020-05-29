package router

import (
	"github.com/gin-gonic/gin"

	"goweb/controller"
)

func router(route *gin.Engine) *gin.Engine {
	route.GET("/hello", controller.TestReturn.Hello)

	return route
}
