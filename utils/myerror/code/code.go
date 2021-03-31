package code

import (
	"fmt"

	"goweb/utils/callback"
)

// ToInt 获取响应码
func ToInt(c callback.ResponseType) int {
	return int(c)
}

// Msg 获取信息
func Msg(c callback.ResponseType) string {
	if v, ok := message[c]; ok {
		return v
	}

	return ""
}

// MsgWithCode 获取错误码+信息 [10001]操作失败，类型不正确
func MsgWithCode(c callback.ResponseType) string {

	if v, ok := message[c]; ok {
		return fmt.Sprintf("[%d]%s", c, v)
	}

	return "[90000]未知错误"
}

// 错误预设+追加
func MsgAppend(c callback.ResponseType, appendStr string) string {
	if v, ok := message[c]; ok {
		return v + appendStr
	}

	return appendStr
}
