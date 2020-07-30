package services

import (
	"github.com/AlexShukhman/zeitns_http/types"
	"github.com/gomodule/redigo/redis"
)

// GetToken - Get ZeitToken from Key
func GetToken(client redis.Conn, key string) (*types.ZeitToken, error) {
	var tok types.ZeitToken

	vals, err := redis.Values(client.Do("HGETALL", redis.Args{}.Add(key)...))

	if err != nil {
		token := types.ZeitToken{}
		return &token, err
	}

	err = redis.ScanStruct(vals, tok)

	if err != nil {
		token := types.ZeitToken{}
		return &token, err
	}

	return &tok, err
}

// CheckKey - Check if key exists
func CheckKey(client redis.Conn, key string) (exists bool, err error) {
	exists, err = redis.Bool(client.Do("EXISTS", redis.Args{}.Add(key)...))

	return exists, err
}
