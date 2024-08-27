package message_service

import (
	"G1/message-service/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/messages", handlers.CreateMessage).Methods("POST")

	log.Println("Starting message service on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
