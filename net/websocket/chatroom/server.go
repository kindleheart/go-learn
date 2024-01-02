package main

import (
	"fmt"
	"goLearn/net/websocket/chatroom"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	go chatroom.h.run()
	router.HandleFunc("/ws", chatroom.myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
