package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// func h(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "hello wisej")
// }

func ws(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}
		log.Println(string(msg))
	}

}
func main() {
	http.HandleFunc("/", ws)
	http.ListenAndServe(":8080", nil)
}
