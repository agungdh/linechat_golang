package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader for WebSocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (customize for security)
	},
}

// Client connections map
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

// Handle WebSocket connections
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	// Register client
	clients[conn] = true
	log.Println("New WebSocket client connected")

	// Listen for incoming messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket Read Error:", err)
			delete(clients, conn)
			break
		}

		log.Println("Received:", string(msg))
		broadcast <- string(msg)
	}
}

// Broadcast messages to all clients
func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("WebSocket Write Error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// Start WebSocket server
func StartServer() {
	http.HandleFunc("/ws", HandleConnections)
	go HandleMessages()

	log.Println("WebSocket server started on ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
