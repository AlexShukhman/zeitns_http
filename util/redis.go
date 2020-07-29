package util

import (
	"github.com/AlexShukhman/zeitns_http/types/Standard"
	"github.com/go-redis/redis"
)

// ConnectRedis - connect to the Redis instance
func ConnectRedis(host string, port uint16, pw string, db uint8) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + string(port),
		Password: pw,
		DB:       int(db),
	})

	var ex *Standard.ZeitToken

	return client
}
