package errors

import (
	"goweb/errors/code"
	"goweb/helper/api"
)

/***** 90000,90500 用下面的方法 *****/

// New APIError API普通错误提示，90000
func New(msg string, code ...api.ResponseType) (ae *api.APIError) {
	c := api.ResponseError
	if len(code) == 1 {
		c = code[0]
	}
	return api.NewAPIError(msg, c)
}

// Except APIError API异常错误提示 90500
func Except(e error) (ae *api.APIError) {
	return api.NewAPIErrorWithLog("系统错误", e.Error())
}

/***** 自定义错误码用下面的方法 *****/

// LogicMask 自定义逻辑错误码 code：90000 msg: code对应的信息
func LogicMask(c api.ResponseType) (ae *api.APIError) {
	return api.NewAPIError(code.Msg(c), api.ResponseError)
}

// LogicMaskrWitdhCode 自定义逻辑错误码 code：90000 msg: [code]+code对应的信息
func LogicMaskrWitdhCode(c api.ResponseType) (ae *api.APIError) {
	// log部分待写
	return api.NewAPIError(code.MsgWithCode(c), api.ResponseError)
}

// Logic 自定义逻辑错误码 code：自定义错误码 msg: code对应的信息
func Logic(c api.ResponseType) (ae *api.APIError) {
	return api.NewAPIError(code.Msg(c), c)
}

// LogicWithCode 自定义逻辑错误码 code：自定义错误码 msg: [code]+code对应的信息
func LogicWithCode(c api.ResponseType) (ae *api.APIError) {
	return api.NewAPIError(code.MsgWithCode(c), c)
}

// LoginAdd 错误信息追加 "添加出错" +err
func LoginAdd(c api.ResponseType, appendString string) (ae *api.APIError) {
	return api.NewAPIError(code.MsgAppend(c, appendString), c)
}
