package main

import (
	"goweb/conf"
	_ "goweb/docs"
	"goweb/routers"
	"goweb/service"
	"goweb/utils/vali"
)

// @title Swagger 生成文档测试
// @version 1.0
// @description 这是一个Swagger文档生成测试
// @host localhost:8088
func main() {
	myDBConf := conf.Main.Mysql
	service.Init(myDBConf)
	routers.Init()
	vali.Init()
}
