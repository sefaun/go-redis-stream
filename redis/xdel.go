package redis

import "go-redis-stream/connections"

func XDel(stream string, ids []string) {

	for _, v := range ids {
		connections.RedisConnection.XDel(stream, v)
	}

}
