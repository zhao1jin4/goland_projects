package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	/*
	   ########通道相关：
	   func (v Value) Send(x reflect.Value)// 发送数据（会阻塞），v 值必须是可写通道。

	   func (v Value) Recv() (x reflect.Value, ok bool) // 接收数据（会阻塞），v 值必须是可读通道。

	   func (v Value) TrySend(x reflect.Value) bool // 尝试发送数据（不会阻塞），v 值必须是可写通道。

	   func (v Value) TryRecv() (x reflect.Value, ok bool) // 尝试接收数据（不会阻塞），v 值必须是可读通道。

	   func (v Value) Close() // 关闭通道
	*/
	//-----channel反射

	all := make(chan int)
	go writeData(all, 20)
	go readData(all)

	time.Sleep(30 * time.Second)
}
func readData(readChan <-chan int) { //只能读，不能写
	fmt.Println("准备读")
	refReadChan := reflect.ValueOf(readChan)
	if read, ok := refReadChan.Recv(); ok { //Recv
		fmt.Println("读到了", read)
	}
	fmt.Println("读到结束")
}
func writeData(writeChan chan<- int, num int) { //只能写，不能读
	fmt.Println("准备写")
	refWriteChan := reflect.ValueOf(writeChan)
	refWriteChan.Send(reflect.ValueOf(num)) //Send
	fmt.Println("写了", num)
}
