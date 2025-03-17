package main

import (
	"linechat/websocket"
	"log"
)

func main() {
	// Start WebSocket Server
	go websocket.StartServer()

	// Start your LINE bot logic here
	log.Println("LINE chatbot running...")

	select {} // Keep the main function alive
}
