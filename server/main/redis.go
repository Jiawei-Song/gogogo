package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

// initPool address连接地址，maxIdle最大空闲连接数, idleTimeout最大空闲时间
func initPool(address string, maxIdle int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:         maxIdle,
		MaxActive:       0,
		MaxConnLifetime: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
