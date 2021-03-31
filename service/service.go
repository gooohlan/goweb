package service

import (
	"goweb/conf"
	mydb "goweb/db"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init(mysqlConf conf.Mysql) {
	db, err = mydb.NewDatabase(&mysqlConf)

}
