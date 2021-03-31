package code

import (
	"goweb/utils/callback"
)

var message map[callback.ResponseType]string

func init() {
	message = make(map[callback.ResponseType]string)
	message[GENERAL_DATA_ADD_HAD] = "数据操作失败，记录已存在"
	message[GENERAL_DATA_NOT_CHANGE] = "数据操作失败，记录未更改"
	message[USER_LOGIN_NOT_EXIST] = "登录失败，用户不存在"
	message[USER_LOGIN_PWD_WRONG] = "登录失败，密码错误"
	message[USER_BUY_NOT_BALANCE] = "购买失败，余额不足"
	message[USER_JWT_PARSE_FAIL] = "Token解析失败，HS256的token解析错误，err:"
	message[USER_JWT_ILLEGAL] = "Token解析失败，Token非法"
	// message[DEMO_DATA_Id_NIL] = "数据操作失败，记录不存在"

	// todo
}

/* 自定义错误码

1、必须1开头 1xxxx
2、常量名可以参考 [模块]_[功能]_[函数]_[detail],非固定格式，按实际业务需求定义。

*/
const (
	// DEMO_DATA_ADD_Had 添加数据数据，记录已存在 （参考第2-3位 01 代表对应的功能）
	GENERAL_DATA_ADD_HAD    callback.ResponseType = 10100
	GENERAL_DATA_NOT_CHANGE callback.ResponseType = 10101
	USER_LOGIN_NOT_EXIST    callback.ResponseType = 10200
	USER_LOGIN_PWD_WRONG    callback.ResponseType = 10201
	USER_JWT_PARSE_FAIL     callback.ResponseType = 10202
	USER_JWT_ILLEGAL        callback.ResponseType = 10204
	USER_BUY_NOT_BALANCE    callback.ResponseType = 10205
	// DEMO_DATA_Id_NIL  apiconstant.ResponseType = 10101

	// to do
)
