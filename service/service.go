package service

import (
    "github.com/go-redis/redis/v8"
    "goweb/conf"
    mydb "goweb/db"
    
    "gorm.io/gorm"
)

var (
    db  *gorm.DB
    rdb *redis.Client
    err error
)

func Init(conf conf.Database) {
    db, err = mydb.NewMysql(&conf.Mysql)
    rdb1, cleanup := mydb.NewRedis(&conf.Redis)
    if err != nil {
        cleanup()
        panic(err)
    }
    rdb = rdb1
}
