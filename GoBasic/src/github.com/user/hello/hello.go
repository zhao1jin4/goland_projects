package main

import (
	f "fmt" //导入别名

	stringutil1 "github.com/user/stringutil" //就可引用刚刚的库,是文件夹名,应该放在 $GOPATH/src/github.com/user/stringutil
	// _ "github.com/user/stringutil" //表示只为执行包中的init函数
)

func main() { //这个括号不能换行
	f.Println(stringutil1.Reverse("!oG ,olleH")) //引用的是包名

	var p = stringutil1.Persion{Name: "lisi"} //只能键传值
	f.Printf("name=%s\n ", p.Name)
	//f.Println(" age=%d",p.age)//age访问不到
}
func init() {
	f.Println("main包Init函数初始") //在main方法前执行
}
