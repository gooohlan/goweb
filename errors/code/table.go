package code

import (
	"grab_parking/helper/api"
)


var message map[api.ResponseType]string

func init() {
	message = make(map[api.ResponseType]string)
	message[DEMO_DATA_ADD_Had] = "数据操作失败，记录已存在"
	// message[DEMO_DATA_ID_NIL] = "数据操作失败，记录不存在"

	// todo
}

/* 自定义错误码

1、必须1开头 1xxxx
2、常量名可以参考 [模块]_[功能]_[函数]_[detail],非固定格式，按实际业务需求定义。

*/
const (
	// DEMO_DATA_ADD_Had 添加数据数据，记录已存在 （参考第2-3位 01 代表对应的功能）
	DEMO_DATA_ADD_Had api.ResponseType = 10100
	// DEMO_DATA_ID_NIL  apiconstant.ResponseType = 10101

	//to do
)

	