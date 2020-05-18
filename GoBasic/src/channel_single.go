package main

import (
	"fmt"
	"time"
)

func main() {
	all := make(chan int )
	go readData(all)
	go writeData(all,10)

	//以下代码没什么作用
	writeChan := make(chan <- int )//只能写，不能读
	readChan := make(<- chan  int )//只能读，不能写
	go writeData(writeChan,20)
	go readData(readChan)

	time.Sleep(30*time.Second)
}
func readData(readChan <- chan  int){//只能读，不能写
	fmt.Println("准备读")
	read:= <-readChan //如是单向的这个执行后，不会执行后面的代码
	//readChan <-  1//报错
	fmt.Println("读到了",read)
}
func writeData(writeChan chan <-  int,num int){//只能写，不能读
	fmt.Println("准备写")
	writeChan <-  num //如是单向的这个执行后，不会执行后面的代码
	//readErr:= <- writeChan//报错
	fmt.Println("写了",num)
}