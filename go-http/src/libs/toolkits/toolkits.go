package toolkits

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

/*
golang interface -> json string without escape

params:
	v|interface

return:
	string
*/
func InterfaceToJsonStr(v interface{}) string {
	var buffer = bytes.NewBufferString("")

	var enc = json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	var _ = /*error ignore*/ enc.Encode(v)

	return buffer.String()
}

/*
json string -> golang type

params:
	js|string
	v|interface

return:
	error
*/
func JsonStrToInterface(js string, v interface{}) error {
	var dec = json.NewDecoder(bytes.NewBufferString(js))
	var err = dec.Decode(v)

	return err
}

/*
random number with seed(unix nano)
*/
func Random(n int) int {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var num = r.Intn(n)
	return num
}

/*
compare a with b, return the max number
*/
func CompareWithN(a int) func(int) int {
	var fn = func(b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	return fn
}

/*
get the name of type
*/
func Type(v interface{}) (typ string) {
	switch vt := v.(type) {
	default:
		//val = fmt.Sprintf("%s", vt)
		typ = fmt.Sprintf("%T", vt)
	}
	return
}
