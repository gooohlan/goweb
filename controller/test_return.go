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

func (t *testReturn) Registered(c *gin.Context) {
	type User struct {
		Name string
		Age int
	}
	qry := new(User)
	// 使用验证器验证数据格式
	if err := API.ParseRequest(c, qry); err != nil {
		return
	}
	return
}
