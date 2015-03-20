//模型层基础
package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//初始化内容
func init() {

	//创建引擎
	engine, err = xorm.NewEngine("sqlite3", "./erp.db")

	//数据库同步

}

var engine *xorm.Engine
