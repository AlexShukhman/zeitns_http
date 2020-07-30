package main

import (
	"log"

	"github.com/AlexShukhman/zeitns_http/util"
)

func main() {
	conn := util.ConnectRedis("localhost", 6379, "")
	defer conn.Close()

	_, err := conn.Do("PING")

	if err != nil {
		log.Fatal(err)
	}
}
