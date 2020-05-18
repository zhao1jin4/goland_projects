package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)//等待
		fmt.Println(s)
	}
}

func main() {
	//协程coroutine(协同程序),GO语言命名为Goroutine
	//main函数叫主Goroutine ,如主Goroutine结束，所有子的Goroutine也被中止了


	//每一个用户线程 和内核线程是一对一的关系 （Java/C++使用的方式）
	//多个用户线程 对应 一个内核线程(只能用一个核，但减少切换开销)，是一对多的关系 ,如一个线程在阻塞中，其它所有线程都不会被调度到，修改为非阻塞式库，在要阻塞前让出CPU，通知其它用户线程
	//多个用户线程 对应 多个内核线程 是多对多的关系，(减少用户内核切换开销)可在运行时动态关联，当一个内核线程上的一个用户线程阻塞，这个内核线程上的其它用户线程，会被调试到其它内核线程上 (GO使用，自己实现了一个运行调试器)
	/*schedular实现有4个结构
	 Sched
	Machine(由操作系统管理的，线程，用来运行Goroutine)
	Processor 维护了Goroutine队列,从N:1,到N:M，个会申请批量资源
	Goroutine 如需要资源先向Processor申请

	CSP模型=communicating Sequential Process ,GO语言的CSP是用gorouting和channel实现
	channel 先进先出（队列）,底层也是用mutex，只是功能更强，类似  unix 的pipe

	*/
	go say("world") //启动新协程，函数有返回值也会被忽略
	say("hello")
}