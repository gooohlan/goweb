package myerror

import (
	"goweb/utils/callback"
	"goweb/utils/myerror/code"
)

/***** 90000,90500 用下面的方法 *****/

// New APIError API普通错误提示，90000
func New(msg string, code ...callback.ResponseType) (ae *callback.APIError) {
	c := callback.ResponseError
	if len(code) == 1 {
		c = code[0]
	}
	return callback.NewAPIError(msg, c)
}

// Except APIError API异常错误提示 90500
func Except(e error) (ae *callback.APIError) {
	return callback.NewAPIErrorWithLog("系统错误", e.Error())
}

/***** 自定义错误码用下面的方法 *****/

// LogicMask 自定义逻辑错误码 code：90000 msg: code对应的信息
func LogicMask(c callback.ResponseType) (ae *callback.APIError) {
	return callback.NewAPIError(code.Msg(c), callback.ResponseError)
}

// LogicMaskrWitdhCode 自定义逻辑错误码 code：90000 msg: [code]+code对应的信息
func LogicMaskrWitdhCode(c callback.ResponseType) (ae *callback.APIError) {
	// log部分待写
	return callback.NewAPIError(code.MsgWithCode(c), callback.ResponseError)
}

// Logic 自定义逻辑错误码 code：自定义错误码 msg: code对应的信息
func Logic(c callback.ResponseType) (ae *callback.APIError) {
	return callback.NewAPIError(code.Msg(c), c)
}

// LogicWithCode 自定义逻辑错误码 code：自定义错误码 msg: [code]+code对应的信息
func LogicWithCode(c callback.ResponseType) (ae *callback.APIError) {
	return callback.NewAPIError(code.MsgWithCode(c), c)
}

// LocallbackAdd 错误信息追加 "添加出错" +err
func LocallbackAdd(c callback.ResponseType, appendString string) (ae *callback.APIError) {
	return callback.NewAPIError(code.MsgAppend(c, appendString), c)
}
