package controller

import (
	"github.com/gin-gonic/gin"
	
	"grab_parking/errors/code"
	"grab_parking/helper/api"
	"grab_parking/errors"
)

var TestReturn testReturn
type testReturn struct {

}

var API api.API

func(t * testReturn) Hello(c *gin.Context)  {
	err := errors.LogicWithCode	(code.DEMO_DATA_ADD_Had)
	API.SetError(c, err)
}
