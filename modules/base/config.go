//config
package base

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	AppName     string //app名称
	Port        int    //端口
	Domain      string //域名
	CrossDomain bool   //是否支持跨域请求
	DbUrl       string //数据库连接地址
	DbUsername  string //用户名
	DbPassword  string //密码
}

const (
	//配置文件名称
	CONFIG_NAME = "app.conf"
)

var conf Config

//默认加载配置内容
func init() {
	file, err := os.Open(CONFIG_NAME)
	if err != nil {
		log.Fatalln(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetConfig() *Config {
	return &conf
}
