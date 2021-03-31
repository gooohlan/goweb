package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"goweb/controller"
)

func Init() {
	route := gin.Default()
	route = router(route)
	_ = route.Run(":8088")
}

func router(route *gin.Engine) *gin.Engine {
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	tagRoute := route.Group("tag")
	tagRoute.GET("", controller.TestReturn.Hello)
	tagRoute.POST("", controller.TestReturn.Hello)
	tagRoute.PUT("/:id", controller.TestReturn.Hello)
	tagRoute.DELETE("/:id", controller.TestReturn.Hello)
	return route
}
