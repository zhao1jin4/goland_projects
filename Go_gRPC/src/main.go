package main

/*

有负载均衡功能,跟踪，监控，身份验证

https://github.com/grpc/grpc-go
go get google.golang.org/grpc 下不了就

go get -u github.com/grpc/grpc-go 或  下载下来 放到 $GOPATH/sc 下修改目录为google.golang.org/grpc

还依赖于 golang.org/x/net  就是  https://github.com/golang/net  下载master很快，再移动目录
		golang.org/x/text  就是  https://github.com/golang/text 下载master很慢

		google.golang.org/genproto 就是 https://github.com/googleapis/go-genproto 有时会多次失败
		git clone https://github.com/googleapis/go-genproto
		go get  github.com/googleapis/go-genproto  再移动目录

如报 runnerw.exe: CreateProcess failed with error 216: This version of %1 is not compatible with the version of Windows you're running. Check your computer's system information to see whether you need a x86 (32-bit) or x64 (64-bit) version of the program, and then contact the software publisher.
是因为main方法所在文件的包名不是main包

protoc -I=. --go_out=. hello.proto  生成的hello.pb.go文件，没有 service 声明的方法SayHello
protoc -I=. --go_out=plugins=grpc:. hello.proto  就有service 方法了,这个和示例代码中的还是有区别的
*/

func main(){


}

