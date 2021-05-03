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
	会自动下载 https://github.com/golang/protobuf 项目  放 github.com/golang/protobuf 也可手动做


cd  $GOPATH/src/github.com/golang/protobuf/protoc-gen-go
go install 生成了 protoc-gen-go.exe 在$GOPATH/bin目录下
			  自动 依赖 生成/下载 了 $GOPATH\pkg\mod\google.golang.org\protobuf@v1.23.0  可使用这个目录 复制到 src目录下(去版本号)
			  也可事先手工单独下载ttps://github.com/protocolbuffers/protobuf-go 项目,放在 ~\go\src\google.golang.org\protobuf


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
