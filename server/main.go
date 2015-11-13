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
	log.Println("New Connection")
	for {
		messageType, messageBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Connection Closed.")
			break
		}
		if messageType != websocket.BinaryMessage {
			log.Println("Message not binary:\n", string(messageBytes))
			break
		}

		log.Println("Received Message:\n", string(messageBytes))
	}
	conn.Close()
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
	log.Println("Server Listening on :3000")
	defer log.Println("Server done.")
	http.ListenAndServe(":3000", nil)
}
