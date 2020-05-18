package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex  //也有RWMutext,也能检查到死锁(如有时候没有UnLock)，比java牛
	//不要共享内存的方式去通讯，而是以通讯的方式共享内存,应该用channel
	var global int=10;
	go func(){
		//go run -race bad_lock.go 如不加锁，会显示警告
		var isDone bool=false;
		for !isDone{
			mutex.Lock()
			if global>0 {
				time.Sleep(10*time.Millisecond)
				global--;//可以仿问到外部的变量
				fmt.Println("No Name库存少了一个",global)
			}else {
				isDone=true;
			}
			mutex.Unlock()
		}
	}()

	var isDone bool=false;
	for !isDone{
		mutex.Lock()
		if global>0 {
			time.Sleep(10*time.Millisecond)
			global--;//可以仿问到外部的变量
			fmt.Println("main库存少了一个",global)
		}else {
			isDone=true
		}
		mutex.Unlock()
	}
	time.Sleep(2*time.Second)
}