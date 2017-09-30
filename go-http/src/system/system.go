package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
)

/*
 加载json配置文件
*/
func LoadConfigure(path string) (map[string]interface{}, error) {
	var b, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var m = make(map[string]interface{})

	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index")
	return
}

/*
 启动HTTP服务
*/
func StartHttpServer() {
	//1
	//var server_mux = http.NewServeMux()

	//2
	//var mux = routes.NewMux()
	//var handler, _ = mux.BuildRoutes([]routes.Handler{
	//	mux.Handler("GET", "/index", Index),
	//})
	//var server = &http.Server{
	//	Addr:    ":4242",
	//	Handler: mux,
	//}

	//http.ListenAndServe(":4242", nil)
	//or
	var server = &http.Server{
		Addr: ":4242",
		//Handler: server_mux,
		Handler: nil,
	}

	var _ = server.ListenAndServe()

	//var mux = denco.NewMux()
	//handler, err := mux.Build([]denco.Handler{
	//	mux.GET("/你好", Index),
	//})
	//if err != nil {
	//	panic(err)
	//}
	//http.ListenAndServe(":4242", handler)
	return
}

/*TEST AREA -----------------------------------------------------------------*/
type counter struct {
	number int
}

func (c *counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.number++
	fmt.Fprintf(w, "counter = %d", c.number)
}
