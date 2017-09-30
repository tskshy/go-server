package routes

//import (
//	"fmt"
//	"net/http"
//)
//
//func NewServeMux() *ServeMux {
//	return &ServeMux{
//		Routers: map[string]*Records{},
//	}
//}
//
//type ServeMux struct {
//	Routers map[string][]*Records
//}
//
//type Records struct {
//	/*path pattern*/
//	Pattern string
//	/*handler function*/
//	Func HandlerFunc
//}
//
//type HandlerFunc func(w http.ResponseWriter, r *http.Request)
//
//func (mux *ServeMux) BuildRoutes(routers []TODOName) {
//	for _, item := range routers {
//		fmt.Println(item)
//	}
//}
//
//type TODOName struct {
//	Method string
//	Path   string
//	Func   HandlerFunc
//}
//
//func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	var router, found = mux.Routers[r.Method]
//
//	if found && router.Pattern == r.URL.Path {
//		router.Func(w, r)
//		return
//	}
//
//	NotFound(w, r)
//	return
//}
//
//var NotFound = func(w http.ResponseWriter, r *http.Request) {
//	http.NotFound(w, r)
//}
//
//func Analyze(path string) (HandlerFunc, string, bool) {
//	return nil, "", false
//}
