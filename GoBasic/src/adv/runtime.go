package main

import (
	"fmt"
	"runtime"
	"time"
)
func init(){
	runtime.GOMAXPROCS(8)
	fmt.Println("CPU核数=",runtime.NumCPU())
}
func main() {
	 //runtim相当于JVM 编译出来的程序是运行在runtime上，负责内存分配，拉圾回收，反射，goroutine,channel
	 fmt.Println("GOROOT="+runtime.GOROOT())

	fmt.Println("OS="+runtime.GOOS)//windows,darwin

	go func(){
	 	for i:=0;i<10;i++{
			fmt.Println(i)
		}
	}()
	runtime.Gosched()//让出CPU

	go myfunc();


	time.Sleep(3*time.Second)
}
func myfunc(){
	defer  fmt.Println("myfunc defer")
	//return //后面有代码不会报错
	runtime.Goexit()//退出goroutine,所在方法必须是用go启动的,还会执行defer
	fmt.Println("myfunc done")
}

