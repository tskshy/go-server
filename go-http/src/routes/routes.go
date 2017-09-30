package routes

import (
	"fmt"
	"github.com/naoina/denco"
	"net/http"
)

/*
 http://www.cnblogs.com/hitandrew/p/5816925.html

 http://blog.csdn.net/xingwangc2014/article/details/51623157
*/

func Controller() (http.Handler, error) {
	denco.NotFound = func(w http.ResponseWriter, r *http.Request, _ denco.Params) {
		//fmt.Fprintf(w, "404 404")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	var mux *denco.Mux = denco.NewMux()

	var handler, err = mux.Build([]denco.Handler{
		mux.GET("/:name", func(w http.ResponseWriter, r *http.Request, params denco.Params) {
			var num = 4 / 1
			fmt.Fprintf(w, "hello %s %d", params.Get("name"), num)
		}),
	})

	return handler, err

}
