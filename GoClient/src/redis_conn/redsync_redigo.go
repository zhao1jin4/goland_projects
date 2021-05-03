package main

import (
	"time"

	"github.com/go-redsync/redsync"
	"github.com/go-redsync/redsync/redis/redigo"
	"github.com/gomodule/redigo/redis"
)

//redlock的实现 redigo ,测试启两个，都可以得到锁？？？
func main() {
	setdb := redis.DialDatabase(2)
	setPasswd := redis.DialPassword("123")
	var innerPool *redis.Pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", ":6379", setdb, setPasswd) },
	}

	pool := redigo.NewPool(innerPool)
	rs := redsync.New(pool)
	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	mutexname := "my-global-mutex"
	mutex := rs.NewMutex(mutexname)

	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := mutex.Lock(); err != nil {
		panic(err)
	}

	// Do your work that requires the lock.

	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}
}
