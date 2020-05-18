package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int,done chan bool) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		fmt.Println("写了",x)
		x, y = y, x+y
	}
	close(c)//关闭channel,表示没有更多的数据了
	done<-true
}

func main() {
	c := make(chan int,10)//缓冲区,可以写<=10个，不会阻塞
	done := make(chan bool)
	go fibonacci(cap(c), c,done) //slice 的cap函数 容量

	//如果上面的 c 通道不关闭，那么 range 函数就报错
	//for i := range c { //range 来(遍历)取 channel,是在关闭之后退出
	//	fmt.Println(i)
	//}
	//--方式二
	for {
		time.Sleep(100*time.Millisecond)
		v,ok := <- c
		fmt.Printf("len=%d,cap=%d,ok=%t\n",len(c),cap(c),ok)
		if ok {//true表示没有关闭
			fmt.Println("读了",v)
		}else {
			break;
		}
	}
	//-----阻塞方式 ，多传一个chan,
	<- done //可以不接收变量，这种方式阻塞，只会传一个值过来
}