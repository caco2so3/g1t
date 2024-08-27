package redis

import (
	"G1/message-service/models"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

func PublishMessage(msg models.Message) {
	data, _ := json.Marshal(msg)
	rdb.Publish(ctx, "messages_channel", data)
}
