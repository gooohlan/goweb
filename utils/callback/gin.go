package callback

import (
	"encoding/json"
	"fmt"
	"net/http"

	"goweb/utils/vali"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CallbackData struct {
	ctx       *gin.Context
	result    ApiResult
	rawResult []byte
}

type ApiResult struct {
	Code   ResponseType `json:"code"`
	Msg    string       `json:"msg"`
	Data   interface{}  `json:"data"`
	dataKV map[string]interface{}
}

// Return API信息返回
func (a ApiResult) Return(c *gin.Context) {
	if a.Data == nil {
		a.Data = struct{}{}
	}
	c.JSON(http.StatusOK, a)
}

// SetError 设置错误信息
func (a ApiResult) SetError(c *gin.Context, err error) {
	a.Msg = err.Error()
	a.Code = ResponseError
	if e, ok := err.(*APIError); ok {
		a.Code = e.code
		a.Data = e.data
	}
	a.Return(c)
}

// SetMsg 设置信息，code默认ResponeError
func (a ApiResult) SetMsg(c *gin.Context, msg string, code ...ResponseType) {
	a.Msg = msg
	a.Code = ResponseError
	if len(code) == 1 {
		a.Code = code[0]
		a.Return(c)
	}
	a.Return(c)
}

// SetData 设置输出的model
func (a ApiResult) SetData(c *gin.Context, data interface{}) {
	a.Code = ResponseOK
	a.Data = data
	a.Return(c)
}

// 空返回时使用
func (a ApiResult) SetEmpty(c *gin.Context) {
	a.Code = ResponseOK
	a.Return(c)
}

// SetDataKV 设置KV，会覆盖掉 SetData
func (a ApiResult) SetDataKV(c *gin.Context, key string, value interface{}) {
	a.Code = ResponseOK
	if a.dataKV == nil {
		a.dataKV = make(map[string]interface{})
	}
	a.dataKV[key] = value
	a.Data = a.dataKV
	a.Return(c)
}

func (a ApiResult) ParseRequest(c *gin.Context, request interface{}) error {
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
		a.Code = ResponseRequest
		a.Msg = errStr
		a.Return(c)
		return err
	}
	return nil
}
