package util

import (
	"log"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

// ConnectRedis - connect to the Redis instance
func ConnectRedis(host string, port uint16, pw string) redis.Conn {
	c, err := redis.Dial(
		"tcp",
		host+":"+strconv.Itoa(int(port)),
		redis.DialKeepAlive(1*time.Second),
		redis.DialPassword(pw),
		redis.DialConnectTimeout(5*time.Second),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	return c
}
