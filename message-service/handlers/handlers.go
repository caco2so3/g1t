package handlers

import (
	"G1/message-service/models"
	"G1/message-service/redis"
	"G1/message-service/repository"
	"encoding/json"
	"net/http"
	"time"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(msg.Content) > 128 {
		http.Error(w, "Message exceeds 128 characters", http.StatusBadRequest)
		return
	}

	msg.Timestamp = time.Now()

	err = repository.SaveMessage(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	redis.PublishMessage(msg)

	w.WriteHeader(http.StatusCreated)
}
