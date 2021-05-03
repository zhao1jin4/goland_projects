package main

//go mod init github.com/my/repo 这个是测试
//go get github.com/go-redis/redis/v8
import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Password: "", // no password set
		Password: "123",
		DB:       0, // use default DB
	})
	{
		err := rdb.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", val)

		val2, err := rdb.Get(ctx, "key2").Result()
		if err == redis.Nil {
			fmt.Println("key2 does not exist")
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("key2", val2)
		}
		// Output: key value
		// key2 does not exist
	}
	{
		// SET key value EX 10 NX
		set, err := rdb.SetNX(ctx, "key", "value", 10*time.Second).Result()
		myPrint(set, err)

		//-- SET key value keepttl NX,   err为ERR syntax error ？？
		// set1, err := rdb.SetNX(ctx, "key", "value", redis.KeepTTL).Result() //KeepTTL=-1
		// myPrint(set1, err)

		// SORT list LIMIT 0 2 ASC
		vals, err := rdb.Sort(ctx, "list", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()
		myPrint(vals, err)

		// ZRANGEBYSCORE zset -inf +inf WITHSCORES LIMIT 0 2
		vals1, err := rdb.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
			Count:  2,
		}).Result()
		myPrint(vals1, err)
		// ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
		vals2, err := rdb.ZInterStore(ctx, "out", &redis.ZStore{
			Keys:    []string{"zset1", "zset2"},
			Weights: []float64{2, 3},
		}).Result()
		myPrint(vals2, err)
		// EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
		vals3, err := rdb.Eval(ctx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()
		myPrint(vals3, err)
		// custom command
		res, err := rdb.Do(ctx, "set", "key", "value").Result()
		myPrint(res, err)
	}
}
func myPrint(res interface{}, err error) {
	if err != nil {
		fmt.Println("error")
		panic(err)
	} else {
		fmt.Println(res)
	}
}
