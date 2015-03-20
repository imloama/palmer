//字典
package sys

//行业
type Industry struct {
	Id     int64    `json:"id"`
	Code   string   `xorm:"varchar(50)";json:"code"`
	Name   string   `xorm:"varchar(50) notnull unique 'name'";json:"name"`
	Parent Industry `json:"parent"`
	Remark string   `xorm:"varchar(200)";json:"remark"`
}
type City struct {
	Id   int64  `json:"id"`
	Name string `xorm:"varchar(50) notnull unique";json:"name"`
	Pid  int64  `json:"pid"`
}
