package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("www.baidu.com")
	if err != nil {
		fmt.Println(err)
		if ins, ok := err.(*net.DNSError); ok {
			fmt.Printf("timeout=%t\n", ins.Timeout())
		}
	}
	fmt.Printf("addr=%s\n", addr)
}
