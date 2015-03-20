//render
package base

import (
	"encoding/json"
	"net/http"
)

//返回Json
func RenderJson(code int, w http.ResponseWriter, r *http.Request, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		//code
	}
	w.Header().Set("content-type", "application/json")
	w.Write(b)
}
