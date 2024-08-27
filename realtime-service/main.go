package realtime_service

import (
	"G1/realtime-service/websocket"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", websocket.HandleConnections)

	go websocket.HandleMessages()

	log.Println("Starting WebSocket server on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
