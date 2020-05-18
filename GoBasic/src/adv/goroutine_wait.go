package main

import (
	"fmt"
	"sync"
	"time"
)

var waitG=sync.WaitGroup{}
func main() {
	waitG.Add(2)//如这里是3,会报死锁，这个比java牛，知道所有的goroutine已经结束但还>0
	go func(){
		for i:=1;i<10;i++{
			time.Sleep(4)
			fmt.Println("no name func i=",i)
		}
		waitG.Done() //源码就是Add(-1)
	}()
	go subFunc()
	waitG.Wait()//主协程 等所有子协程执行完，像Java的CountDownLatch
}
func subFunc(){
	defer  waitG.Done()
	for i:=1;i<10;i++{
		time.Sleep(2)
		fmt.Println("subFunc i=",i)
	}
	fmt.Println("myfunc done")
}
