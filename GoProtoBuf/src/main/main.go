package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"myproto"
)



/*

安装 protoc 安装在PATH中
protoc -I=. --go_out=. person.proto (生成 person.pb.go文件)依赖于 protoc-gen-go 插件
 -IPATH, --proto_path=PATH

go get  github.com/golang/protobuf/protoc-gen-go
翻墙 git install
下载了 protoc-gen-go.exe 在$GOPATH/bin目录下 就是  https://github.com/golang/protobuf/tree/master/protoc-gen-go
下载了 $GOPATH\pkg\mod\google.golang.org\protobuf@v1.22.0

github.com/golang/protobuf 依赖导入 google.golang.org/protobuf/xxx (运行时也找不到)
并且在github上是没有encoding和runtime目录的，依赖如下
"google.golang.org/protobuf/encoding/prototext"
"google.golang.org/protobuf/encoding/protowire"
"google.golang.org/protobuf/runtime/protoimpl"
---方式1
翻墙下载的 $GOPATH\pkg\mod\google.golang.org\protobuf@v1.22.0 放在src目录下，没有版本

----方式2
其实是 https://github.com/protocolbuffers/protobuf-go 项目的源码
go get github.com/protocolbuffers/protobuf-go  下载到 $GOPATH\pkg\mod\github.com\protocolbuffers\protobuf-go@v1.23.0
放在src目录下，没有版本 ,修改包名~\go\src\google.golang.org\protobuf

 */

func main(){
	 myId:=[] int32 { *proto.Int32(101) }
	 msg:=& myproto.Person{
	 	 Name: *proto.String("李四"),
		 Id:  myId,
	 	 Email: *proto.String("aa@bb.com"),
	 }
	 encByte,err:=proto.Marshal(msg) //序列化
	if err!=nil {
		fmt.Print("序列化错误=",err)
		return
	}

	 myPerson:=myproto.Person{}
	 err=proto.Unmarshal(encByte,&myPerson)
	if err!=nil {
		fmt.Print("返序列化错误=",err)
		return
	}
	fmt.Println("myPerson=",myPerson)
	fmt.Println("myPerson.Name=",myPerson.Name)
}
