//主运行程序
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mazhaoyong/palmer/modules/base"
	"github.com/mazhaoyong/palmer/modules/router"
	"log"
	"net/http"
)

func main() {
	//获取配置信息，位于app.conf中
	conf := base.GetConfig()
	fmt.Printf("%v ..", conf)

	r := mux.NewRouter()
	r.HandleFunc("/", router.Index)
	//支持跨域
	handle := base.RecoveryHandler(base.CORSHandler(r))

	//端口
	port := 8080
	if conf != nil && conf.Port > 0 {
		port = conf.Port
	}
	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, handle)
	if err != nil {
		log.Fatal("服务器启动失败！", err)
	} else {
		fmt.Println("服务器启动成功！")
	}

}
