package util

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/AlexShukhman/zeitns_http/types"
	"github.com/go-redis/redis"
)

// ConnectRedis - connect to the Redis instance
func ConnectRedis(host string, port uint16, pw string, db uint8) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(int(port)),
		Password: pw,
		DB:       int(db),
	})

	fmt.Println(client.Options().Addr)

	ping, err := json.Marshal(types.ZeitToken{
		SDP: "this",
	})

	if err != nil {
		log.Fatal(err)
	}

	err = client.Set("ping", ping, 0).Err()

	if err != nil {
		log.Fatal(err)
	}

	val, err := client.Get("ping").Result()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)

	return client
}
