package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	fmt.Printf("channel=%T,%v\n",c,c)//%v显示值内存地址，按内存地址传递
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c,如没有读也是阻塞的，容量就是一个,同一时间无论读写，只有一个能操作
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) //建立channel用来做协程间的通讯, 默认情况下，通道是不带缓冲区的,阻塞式的，可以保证同一时间只有一个goroutine取channel数据
	fmt.Printf("channel=%T,%v\n",c,c)//%v显示值内存地址，按内存地址传递
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	time.Sleep(2*time.Second)//测试写一个值也要等
	x, y := <-c, <-c // 从通道 c 中接收,如没有人写,这里阻塞，像 ArrayBlockingQueue

	fmt.Println(x, y, x+y)

	//c1:= make(chan string)
	//c1 <- "hello" //如只有读或写的一方，就会报死锁错误


}