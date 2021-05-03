package main

import (
	service2 "adv/mynet/service"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type MyService struct {
}

//要用大写方法名,参数只可两个，有一个是指针类型的输出,可以都做成结构体
//(s MyService)或者 (s *MyService)
func (s MyService) Calc(req int, resp *int) error {
	*resp = req + 1
	return nil
}
func main() {
	service := new(MyService)
	//err:=rpc.Register(service) //注册服务，默认服务名 为结构体名MyService
	err := rpc.RegisterName("MyService", service) //指定服务名
	if err != nil {
		fmt.Print("注册失败=", err)
		return
	}

	//--结构体参数
	userService := new(service2.UserService)
	rpc.RegisterName("UserService", userService)

	rpc.HandleHTTP() //使用HTTP协议来访问
	//上下这两段，代码没有联系，但运行确实关联上了
	listen, err := net.Listen("tcp", "0.0.0.0:8181") //端口
	if err != nil {
		fmt.Print("监听失败=", err)
		return
	}
	fmt.Print("服务启动了")
	http.Serve(listen, nil) //启动服务器，指定监听端口
}
