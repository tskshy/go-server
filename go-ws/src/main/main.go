package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var UPGRADER = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ws(w http.ResponseWriter, r *http.Request) {
	//UPGRADER.CheckOrigin = func(r *http.Request) bool { return true }
	var ws_conn, ws_err = UPGRADER.Upgrade(w, r, nil)
	if ws_err != nil {
		fmt.Println(ws_err.Error())
		return
	}

	defer func() { fmt.Println("clear") }()
	defer func() { var _ = ws_conn.Close() }()

	for {
		var message_type, message, r_err = ws_conn.ReadMessage()
		//var message_type, _, r_err = ws_conn.NextReader()
		if r_err != nil {
			fmt.Println(r_err.Error())
			break
		}
		var recv_msg = string(message)

		//fmt.Println("recv:", message_type, string(message))
		var w_err = ws_conn.WriteMessage(message_type, []byte(fmt.Sprintf("%s", recv_msg)))
		if w_err != nil {
			fmt.Println(w_err.Error())
			break
		}

		if string(message) == "quit" {
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", ws)
	var _ = http.ListenAndServe(":4242", nil)
}
