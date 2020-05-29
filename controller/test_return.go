package controller

import (
	"github.com/gin-gonic/gin"

	"goweb/errors"
	"goweb/errors/code"
	"goweb/helper/api"
)

var TestReturn testReturn

type testReturn struct {
}

var API api.API

func (t *testReturn) Hello(c *gin.Context) {
	err := errors.LogicWithCode(code.DEMO_DATA_ADD_Had)
	API.SetError(c, err)
}
