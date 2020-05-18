package main

import (
	"fmt"
	"time"
)
func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)


	go func(){
		time.Sleep(1*time.Second) //两个goroutine等待时间相同1秒，每次执行哪个是不定的
		ch1<-1
	}()
	go func(){
		time.Sleep(1*time.Second)
		ch2<-"abc"
	}()


	select //select会随机执行一个(其它的被忽略了)可运行的case(全是通道)。如果没有case可运行，再看是否有default，有就执行，如没有default将阻塞，直到有case可运行。
	{
	case c1:= <-ch1 :
		fmt.Println("ch1=", c1)
	case c2,ok := <-ch2:
		if ok {
			fmt.Println("ch2", c2)
		}else {
			fmt.Println("ch2 读失败")
		}
	case mytime:= <- time.After(2*time.Second) :
		fmt.Println("mytime=",mytime)
	//default:
	//	fmt.Println("default")
	}

}