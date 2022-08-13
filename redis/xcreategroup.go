package redis

import (
	"go-redis-stream/connections"
	"log"
)

func XCreateGroup(stream string, stream_group string) {

	_, err := connections.RedisConnection.XGroupCreateMkStream(stream, stream_group, "$").Result()

	if err != nil {
		log.Printf("%s", err)
	}

}
