package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex *sync.RWMutex
var wait  *sync.WaitGroup
var global int=10;
func main() {
	mutex =new (sync.RWMutex) //指针类型实例化
	wait=new(sync.WaitGroup)
	wait.Add(4)

	go readData(1)
	go readData(2)
	go writeDate(3)
	go writeDate(4)

	wait.Wait()
}
func readData(goId int){
	var isDone bool=false;
	for !isDone{
		fmt.Println(goId,"readData准备读")
		mutex.RLock()
		if global<=0 {
			isDone=true;
		}
		time.Sleep(10*time.Millisecond)
		fmt.Println(goId,"readData读到了global",global)
		mutex.RUnlock()
	}
	wait.Done()
}
func writeDate(goId int){
	var isDone bool=false;
	for !isDone{
		fmt.Println(goId,"writeDate准备写")
		mutex.Lock()
		if global>0 {
			fmt.Println(goId,"正在写")
			time.Sleep(10*time.Millisecond)
			global--;//可以仿问到外部的变量
			fmt.Println(goId,"writeDate库存少了一个",global)
		}else {
			isDone=true
		}
		mutex.Unlock()
	}
	wait.Done()
}