package redis

import (
	"fmt"
	"go-redis-stream/connections"
	"log"

	"github.com/go-redis/redis"
)

func XReadGroup(stream string, stream_group string, stream_listen_status string) []string {

	entries, err := connections.RedisConnection.XReadGroup(
		&redis.XReadGroupArgs{
			Group:    stream_group,
			Streams:  []string{stream, stream_listen_status},
			Count:    10,
			Block:    0,
			Consumer: "sefa",
			NoAck:    false,
		}).Result()

	if err != nil {
		log.Fatal("Redis Stream Group could not Read ! ", err)
	}

	IDs := []string{}

	for i := 0; i < len(entries[0].Messages); i++ {
		messageID := entries[0].Messages[i].ID
		values := entries[0].Messages[i].Values

		IDs = append(IDs, messageID)
		print(messageID, fmt.Sprintf("Stream ID: %v, Name: %v", values["id"], values["name"]))
	}

	return IDs

}
