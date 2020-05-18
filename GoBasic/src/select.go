package main

import (
	"fmt"
	"time"
)
func Chann(ch chan int, stopCh chan bool) {
	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Second)//暂停一秒
	}
	stopCh <- true
}

func main() {
	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh) //开协程

	for {
		select //select会随机执行一个(其它的被忽略了)可运行的case(全是通道)。如果没有case可运行，再看是否有default，有就执行，如没有default将阻塞，直到有case可运行。
		{ //这个可以换行
		case c = <-ch:
			fmt.Println("Recvice", c)
			fmt.Println("channel")
		case s := <-ch:  //:表示一个新的变量
			fmt.Println("Receive", s)
		case _ = <-stopCh: //9个之前，一直没有值
			goto end
		}
	}
end:
}