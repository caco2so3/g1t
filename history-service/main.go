package history_service

import (
	"G1/history-service/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/messages/history", handlers.GetMessageHistory).Methods("GET")

	log.Println("Starting history service on :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
