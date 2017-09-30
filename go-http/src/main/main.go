package main

import (
	"encoding/json"
	"fmt"
	"hbase"
	"strconv"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	//logs.Debug("this is a test")
	//logs.Info("this is a test")
	//logs.Warn("this is a test")
	//logs.Error("this is a test %s", "sad")

	//system.StartHttpServer()

	/*alg test*/
	//var arr = [][]byte{[]byte("AC"), []byte("ACE"), []byte("BCF")} //, []byte("AD"), []byte("CD"), []byte("CF"), []byte("ZQ")}
	//var trie = algorithm.BuildDFA(arr)

	//for _, v := range trie.Children {
	//	fmt.Println(v)
	//}

	//var conn, _ = redis.Dial("tcp", "127.0.0.1:6379")

	//var data, err2 = conn.Do("RPOP", "test")

	//fmt.Println(data, "|", err2)

	//hbase_test()
	//var n = toolkits.CompareWithN(3)(1)
	//fmt.Println(n)

	//fmt.Println(time.Now().Format("15:04:05"))

	//var t interface{} = map[string]string{"a": "b"}
	//fmt.Println(toolkits.Type(t))

	var data = `[1, 268435498, 268435468, 268435535, 268435536, 268435537, 268435506, 268435480, 268435530]`
	var ret = []uint32{}

	var _ = json.Unmarshal([]byte(data), &ret)
	fmt.Println(ret)

}

func hbase_test() {
	var host = "slave9"
	var port = "9090"

	var table = "ts"
	var family = "f"

	var pf = thrift.NewTBinaryProtocolFactoryDefault()
	var socket, err_socket = thrift.NewTSocket(host + ":" + port)
	if err_socket != nil {
		panic(err_socket)
	}

	var client = hbase.NewTHBaseServiceClientFactory(socket, pf)

	if err := socket.Open(); err != nil {
		panic(err)
	}

	defer func() {
		var _ = socket.Close()
	}()

	/* multiple put*/
	fmt.Println(time.Now().Unix())

	var tput_arr []*hbase.TPut
	for i := 1; i <= 9999; i++ {
		var row_key = strconv.Itoa(i)

		var fill = 4 - len(row_key)

		for i := 0; i < fill; i++ {
			row_key = "0" + row_key
		}

		tput_arr = append(tput_arr, &hbase.TPut{
			Row: []byte(row_key),
			ColumnValues: []*hbase.TColumnValue{
				{
					Family:    []byte(family),
					Qualifier: []byte("time"),
					Value:     []byte("time-" + time.Now().String()),
				},
				{
					Family:    []byte(family),
					Qualifier: []byte("name"),
					Value:     []byte("name-" + row_key),
				},
			},
		})
	}

	var err_pm = client.PutMultiple([]byte(table), tput_arr)
	if err_pm != nil {
		panic(err_pm)
	}
	fmt.Println(time.Now().Format("1970-01-01 00:00:00"))

	/* get
	for i := 0; i < 18; i++ {
		var row_key = strconv.Itoa(i)
		var result, err_r = client.Get([]byte(table), &hbase.TGet{Row: []byte(row_key)})
		if err_r != nil {
			panic(err_r)
		}

		for _, cv := range result.GetColumnValues() {
			fmt.Println("value", string(cv.GetValue()))
			fmt.Println("family", string(cv.GetFamily()))
			fmt.Println("timestamp", cv.GetTimestamp())
			fmt.Println("qualifier", string(cv.GetQualifier()))
		}
	}
	*/

	/* scan
	 */

}
