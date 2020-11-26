package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"goweb/vali"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	ctx       *gin.Context
	result    apiResult
	rawResult []byte
}

type apiResult struct {
	Code   ResponseType `json:"code"`
	Msg    string       `json:"msg"`
	Data   interface{}  `json:"data"`
	dataKV map[string]interface{}
}

// Return API信息返回
func (a API) Return(c *gin.Context) {
	if a.result.Data == nil {
		a.result.Data = struct{}{}
	}
	c.JSON(http.StatusOK, a.result)
}

// SetError 设置错误信息
func (a API) SetError(c *gin.Context, err error) {
	a.result.Msg = err.Error()
	a.result.Code = ResponseError
	if e, ok := err.(*APIError); ok {
		a.result.Code = e.code
		a.result.Data = e.data
	}
	a.Return(c)
}

// SetMsg 设置信息，code默认ResponeError
func (a API) SetMsg(c *gin.Context, msg string, code ...ResponseType) {
	a.result.Msg = msg
	a.result.Code = ResponseError
	if len(code) == 1 {
		a.result.Code = code[0]
		a.Return(c)
	}
	a.Return(c)
}

// SetData 设置输出的model
func (a API) SetData(c *gin.Context, data interface{}) {
	a.result.Code = ResponseOK
	a.result.Data = data
	a.Return(c)
}

// SetDataKV 设置KV，会覆盖掉 SetData
func (a API) SetDataKV(c *gin.Context, key string, value interface{}) {
	a.result.Code = ResponseOK
	if a.result.dataKV == nil {
		a.result.dataKV = make(map[string]interface{})
	}
	a.result.dataKV[key] = value
	a.result.Data = a.result.dataKV
	a.Return(c)
}

// 前端数据验证
func (a API) ParseRequest(c *gin.Context, request interface{}) error {
	err := c.ShouldBind(request)
	var errStr string
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errStr = vali.Translate(err.(validator.ValidationErrors))
		case *json.UnmarshalTypeError:
			unmarshalTypeError := err.(*json.UnmarshalTypeError)
			errStr = fmt.Errorf("%s 类型错误，期望类型 %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String()).Error()
		default:
			errStr = "unknown error."
		}
		a.result.Code = ResponseRequest
		a.result.Msg = errStr
		a.Return(c)
		return err
	}
	return nil
}
