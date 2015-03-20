package sys

import (
	"time"
)

type User struct {
	Id           int64     `json:"id"`
	Name         string    `xorm:"varchar(50) notnull unique 'user_name'";json:"name"`
	Password     string    `xorm:"varchar(100) notnull";json:"password"`
	NickName     string    `xorm:"varchar(50) notnull 'nick_name'";json:"nickname"`
	RegisterTime time.Time `xorm:"varchar(20) created";json:"registertime"` //注册时间
	Gender       bool      `json:"gender"`                                  //性别
	Birthday     time.Time `xorm:"varchar(20) null";json:"birthday"`
	Remark       string    `xorm:"varchar(100) null";json:"remark"` //备注

	IsAdmin bool `json:"isAdmin"` //是否为管理员
}
