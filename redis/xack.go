package redis

import "go-redis-stream/connections"

func XAck(stream string, stream_group string, ids []string) {

	for i := 0; i < len(ids); i++ {
		connections.RedisConnection.XAck(stream, stream_group, ids[i])
	}

}
