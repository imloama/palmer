//index router
package router

import (
	//"fmt"
	"github.com/mazhaoyong/palmer/modules/base"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	base.RenderJson(200, w, r, "Hello,World!")
	//w.Write([]byte("Hello,World!"))
}
