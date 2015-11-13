package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func main() {
	u := url.URL{Scheme: "ws", Host: ":3000", Path: "/"}
	log.Println("Connecting...")

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	c.WriteMessage(websocket.BinaryMessage, []byte("Hello Server! :)"))
	c.Close()
}
