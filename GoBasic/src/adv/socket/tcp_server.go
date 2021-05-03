package main

import (
	"fmt"
	"net"
)

//Go语言将 Non-Block + I/O多路复用 “复杂性”隐藏在Runtime中了
//只需在每个连接对应的goroutine中以“block I/O”的方式对待socket处理即可
func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		conn, err := l.Accept() //会阻塞
		defer conn.Close()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}

		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read error:", err)
		panic(err)
	}
	fmt.Println("readed :", string(buf[0:len]))

}
