//ReturnInfo      server return msg
package base

type ReturnInfo struct {
	Code    int         `json:"code"`    //状态码
	Success bool        `json:"success"` //是否成功
	Data    interface{} `json:"data"`    //返回数据
}
