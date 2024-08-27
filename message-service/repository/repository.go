// repository.go
package repository

import (
	"G1/message-service/models"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password= dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func SaveMessage(msg models.Message) error {
	query := "INSERT INTO messages (content, timestamp, order_id) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, msg.Content, msg.Timestamp, msg.OrderID)
	return err
}
