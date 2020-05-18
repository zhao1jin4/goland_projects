package main

import "fmt"

func main() {
	defer myPrint("main is defer begin")
	myFunc();
	defer myPrint("main is defer end")
	myPrint("main is end")
}
func myPrint(str string){
	fmt.Println(str)
}

func myFunc(){
	defer myPrint("myFunc defer begin ")//相当于 finally
	defer func(){
		if msg:=recover(); msg!=nil { //recover一般用在逻辑前面的defer函数中,相当于catch,上层还会继续执行
			fmt.Println("Recover到的错误消息为",msg)
		}
	}()
	panic("Oh error!")//可传任何类型，给recover函数，后面的代码不被执行，但会执行已经defer的, 相当于throw new Exception
	//不应该出现的网络连不上，文件找不到用error
	//严重的如空指针，下标越界,除0，不应走的代码块，用panic
	fmt.Println("in myFunc")
	defer myPrint("myFunc defer end ")
}
