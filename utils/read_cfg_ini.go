package utils

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

var cfg *goconfig.ConfigFile

func InitConfigIni(filepath string) {
	config, err := goconfig.LoadConfigFile(filepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		panic(err)
	}
	cfg = config
}
