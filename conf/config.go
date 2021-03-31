package conf

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	Database `json:"Database" required:"true"`
}

type Database struct {
	Mysql Mysql `json:"mysql"`
}

type Mysql struct {
	Username string `json:"Username" required:"true"`
	Password string `json:"Password" required:"true"`
	Host     string `json:"Host" default:"127.0.01"`
	Port     string `json:"Port" default:"3306"`
	Charset  string `json:"Charset" default:"utf8mb4"`
	Name     string `json:"Name" required:"true"`
}

var Main = (func() Config {
	var conf Config
	if err := configor.Load(&conf, "/Users/inkdp/Code/study/go/src/goweb/conf/config.ENV.json"); err != nil {
		panic(err.Error())
	}
	return conf
})()
