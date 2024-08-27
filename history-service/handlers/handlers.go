package handlers

import (
	"G1/history-service/repository"
	"encoding/json"
	"net/http"
	"time"
)

func GetMessageHistory(w http.ResponseWriter, r *http.Request) {
	end := time.Now()
	start := end.Add(-10 * time.Minute)

	messages, err := repository.GetMessagesByTimeRange(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}
