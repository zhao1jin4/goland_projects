package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	setdb := redis.DialDatabase(2)
	setPasswd := redis.DialPassword("123")

	var pool *redis.Pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", ":6379", setdb, setPasswd) },
	}
	//使用池

	pool.Close()
}
