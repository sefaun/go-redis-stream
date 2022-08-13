package main

import (
	"go-redis-stream/connections"
	"go-redis-stream/redis"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		print("Error loading .env file")
	}

	connections.CreateRedisConnection()

	const StreamGroup = "SefaGroup"
	const Stream = "SefaStream"
	const recent_data = "0"
	const new_data = ">"
	var Data = map[string]interface{}{
		"id":   1,
		"name": "Sefa Ün",
	}

	//Create Stream Group
	redis.XCreateGroup(Stream, StreamGroup)
	//Send Data to Stream Group
	redis.XAdd(Stream, Data)
	//Get Data From Stream Group
	IDs := redis.XReadGroup(Stream, StreamGroup, new_data)
	//Ack Stream Data
	redis.XAck(Stream, StreamGroup, IDs)
	//Delete Stream Data
	redis.XDel(Stream, IDs)
}
