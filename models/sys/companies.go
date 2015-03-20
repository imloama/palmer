//company
package sys

import (
	"time"
)

//公司信息
type Company struct {
	Id           int64    `json:"id"`
	Code         string   `xorm:"varchar(50)";json:"code"`
	Name         string   `xorm:"varchar(50) notnull unique 'name'";json:"name"`
	Address      string   `xorm:"varchar(100);json:"address"`
	Postcode     string   `xorm:"varchar(6);json:"postcode"`       //邮编
	Industry     Industry `json:"industry"`                        //行业
	MainBusiness string   `xorm:"varchar(100);json:"mainbusiness"` //主营业务
	Telphone     string   `xorm:"varchar(100);json:"telphone"`
	Logo         string   `xorm:"varchar(100);json:"logo"` //公司logo文件地址
	Fax          string   `xorm:"varchar(100);json:"fax"`  //传真

	Remark      string    `xorm:"varchar(200)";json:"remark"` //说明,采用富文本
	DelSymbol   bool      `json:"delsymbol"`                  //删除标识
	UpdatedTime time.Time `xorm:"updated";json:"updatedtime"` //更新时间
	CreatedTime time.Time `xorm:"created";json:"createdtime"` //创建时间
}
