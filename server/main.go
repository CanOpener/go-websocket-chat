package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func connReader(conn *websocket.Conn) {
	defer conn.Close()
	for {
		messageType, messageBytes, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if messageType != websocket.BinaryMessage {
			log.Println("Message not binary: ", string(messageBytes))
			break
		}

		log.Println("Received Message: ", string(messageBytes))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}

	go connReader(conn)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
