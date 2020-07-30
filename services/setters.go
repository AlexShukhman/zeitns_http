package services

import (
	"github.com/AlexShukhman/zeitns_http/types"
	"github.com/gomodule/redigo/redis"
)

// SetFullToken - Save a ZeitToken under the key to Redis
func SetFullToken(conn redis.Conn, key string, token *types.ZeitToken) (err error) {
	_, err = conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(token)...)

	return err
}

// UpdateToken - Update one attribute of a ZeitToken in Redis
func UpdateToken(conn redis.Conn, key string, attr string, val string) (err error) {
	_, err = conn.Do("HSET", redis.Args{}.Add(key).Add(attr).Add(val)...)

	return err
}

// UpdateTokenMany - Update many attributes of a ZeitToken in Redis
func UpdateTokenMany(conn redis.Conn, key string, token *types.ZeitToken) (err error) {
	_, err = conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(token))

	return err
}
