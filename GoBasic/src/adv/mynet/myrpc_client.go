package main

import (
	"adv/mynet/bean"
	"fmt"
	"net/rpc"
	"time"
)

func main(){
	client,err:=rpc.DialHTTP("tcp","127.0.0.1:8181")//连接远程
	if err!=nil {
		fmt.Print("连接失败=",err)
		return
	}
	var req int =20
	var resp *int ;
	//格式注册的  服务名.方法名
	err=client.Call("MyService.Calc",req,&resp)//同步调用远程，响应是指针的指针
	if err!=nil {
		fmt.Print("调用失败=",err)
		return
	}
	fmt.Println("resp=",*resp)//取地址
	//--
	var respAsync *int ; //Go会申请空间
	asynCall:=client.Go("MyService.Calc",req,&respAsync,nil) //异步调用远程
	asynDone:= <- asynCall.Done //读通道
	fmt.Println("asynDone=",asynDone)
	fmt.Println("asynRes=",*respAsync)//取地址

	//--结构体参数
	var userResp *bean.MyResponse //Go会申请空间
	var userReq = bean.MyRequest {10,"lisi",75.4,time.Now()}
	userCall:=client.Go("UserService.QueryUser",userReq,&userResp,nil)
	<-userCall.Done
	fmt.Println("userRes=",*userResp)//取地址
}
