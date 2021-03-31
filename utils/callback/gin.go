package callback

import (
	"encoding/json"
	"fmt"
	"goweb/utils/vali"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CallbackData struct {
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
func (a CallbackData) Return(c *gin.Context) {
	if a.result.Data == nil {
		a.result.Data = struct{}{}
	}
	c.JSON(http.StatusOK, a.result)
}

// SetError 设置错误信息
func (a CallbackData) SetError(c *gin.Context, err error) {
	a.result.Msg = err.Error()
	a.result.Code = ResponseError
	if e, ok := err.(*APIError); ok {
		a.result.Code = e.code
		a.result.Data = e.data
	}
	a.Return(c)
}

// SetMsg 设置信息，code默认ResponeError
func (a CallbackData) SetMsg(c *gin.Context, msg string, code ...ResponseType) {
	a.result.Msg = msg
	a.result.Code = ResponseError
	if len(code) == 1 {
		a.result.Code = code[0]
		a.Return(c)
	}
	a.Return(c)
}

// SetData 设置输出的model
func (a CallbackData) SetData(c *gin.Context, data interface{}) {
	a.result.Code = ResponseOK
	a.result.Data = data
	a.Return(c)
}

// 空返回时使用
func (a CallbackData) SetEmpty(c *gin.Context) {
	a.result.Code = ResponseOK
	a.Return(c)
}

// SetDataKV 设置KV，会覆盖掉 SetData
func (a CallbackData) SetDataKV(c *gin.Context, key string, value interface{}) {
	a.result.Code = ResponseOK
	if a.result.dataKV == nil {
		a.result.dataKV = make(map[string]interface{})
	}
	a.result.dataKV[key] = value
	a.result.Data = a.result.dataKV
	a.Return(c)
}

func (a CallbackData) ParseRequest(c *gin.Context, request interface{}) error {
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
