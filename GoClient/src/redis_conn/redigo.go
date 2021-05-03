package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	setdb := redis.DialDatabase(2)
	setPasswd := redis.DialPassword("123")

	c1, err := redis.Dial("tcp", "127.0.0.1:6379", setdb, setPasswd)
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()
	c2, err := redis.DialURL("redis://127.0.0.1:6379", setdb, setPasswd)
	if err != nil {
		log.Fatalln(err)
	}
	defer c2.Close()

	rec1, err := c1.Do("Get", "name")
	if rec1 != nil {
		fmt.Printf("---------  ")
		fmt.Println(rec1)
	}

	//pipline 管道,redis管道可以用来一次性执行多个命令
	//receive一次只从结果中拿出一个send的命令进行处理
	c2.Send("SET", "foo", "bar")
	c2.Send("GET", "foo")
	c2.Flush()
	v, err := c2.Receive()
	fmt.Println(v) //OK
	v2, err := c2.Receive()
	fmt.Println(string(v2.([]byte)))

	c2.Send("Get", "name")
	c2.Flush()
	rec2, err := c2.Receive()
	if err != nil {
		panic(err)
	}
	if rec2 != nil {
		fmt.Printf("---------  ")
		fmt.Println(string(rec2.([]byte)))
	}

	resset, err := redis.String(c1.Do("SET", "my_test", "redigo"))
	if err != nil {
		fmt.Println("set err")
	} else {
		fmt.Println(resset)
	}

	//获取value并转成字符串
	account_balance, err := redis.String(c1.Do("GET", "my_test"))
	if err != nil {
		fmt.Println("err while getting:", err)
	} else {
		fmt.Println(account_balance)
	}

	//对已有key设置5s过期时间
	n, err := redis.Int64(c1.Do("Expire", "my_test", 5))
	if err != nil {
		fmt.Println(n)
	} else if n != int64(1) {
		fmt.Println("failed")
	}

	n2, err := redis.Int64(c2.Do("TTL", "my_test"))
	if err != nil {
		fmt.Println(n2)
	} else if n != int64(1) {
		fmt.Print(err)
		fmt.Println("failed")
	}

	//删除key
	res2, err := c1.Do("DEL", "my_test")
	if err != nil {
		fmt.Print(err)
		fmt.Println("del err")
	} else {
		fmt.Println(res2)
	}

}
