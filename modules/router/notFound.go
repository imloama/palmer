//NotFound handler
package router

import (
	"net/http"
)

//未找到内容时使用
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
