package main

import (
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	//log.Fatalln("这是一条会触发fatal的日志。")//不会执行后面的代码
	//log.Panicln("这是一条会触发panic的日志。")//不会执行后面的代码

	//如何滚文件呢？
	to, err := os.OpenFile("d:/tmp/go.log", os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	logger := log.New(io.MultiWriter(to, os.Stdout), "<Prodcut_module>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Printf("这是一条%s日志。\n", v)
}
