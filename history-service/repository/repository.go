package repository

import (
	"G1/history-service/models"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password= dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func GetMessagesByTimeRange(start, end time.Time) ([]models.Message, error) {
	query := "SELECT id, content, timestamp, order_id FROM messages WHERE timestamp BETWEEN $1 AND $2"
	rows, err := db.Query(query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.Content, &msg.Timestamp, &msg.OrderID)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
