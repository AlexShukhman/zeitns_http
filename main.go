package main

import (
	"github.com/AlexShukhman/zeitns_http/util"
)

func main() {
	util.ConnectRedis("localhost", 6379, "", 0)
}
