package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	route := gin.Default()
	route = router(route)
	_ = route.Run(":8088")
}
