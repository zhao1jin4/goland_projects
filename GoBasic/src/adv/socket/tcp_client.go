package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", ":8888", 2*time.Second)
	if err != nil {
		fmt.Println("connect error:", err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("你好 from client"))
}
