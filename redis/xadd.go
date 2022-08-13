package redis

import (
	"go-redis-stream/connections"
	"log"

	"github.com/go-redis/redis"
)

func XAdd(stream string, data map[string]interface{}) {

	result, err := connections.RedisConnection.XAdd(&redis.XAddArgs{
		ID:           "",
		Stream:       stream,
		MaxLen:       0,
		MaxLenApprox: 0,
		Values:       data,
	}).Result()

	if err != nil {
		log.Printf("%s", err)
	}

	log.Printf("Data Send To StreamGroup with ID: %s ", result)

}
