//recovery
//遇到panic进行恢复
package base

import (
	"encoding/json"
	"net/http"
)

//问题恢复
func RecoveryHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				//log
				w.Header().Set("content-type", "application/json")
				info := &ReturnInfo{
					Code:    500,
					Success: false,
					Data:    err,
				}
				b, e := json.Marshal(info)
				if e != nil {
					//log
				}
				w.Write(b)

			}
		}()
		handler.ServeHTTP(w, r)
	})

}
