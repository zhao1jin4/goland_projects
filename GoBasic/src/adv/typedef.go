package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyDuration time.Duration
func(my MyDuration)showDur(){ //必须是新的类型才加方法

}

func main(){

	type myint int //相当于C的typedef,是一种新的数据类型
	var num myint=20
	//var num2 int =30
	//num=num2//错误
	fmt.Printf("type=%T,num=%d\n",num,num)

	type mystring=string //是别名,可以向以前版本兼容

	var str mystring="abc";
	var str2 string ="def"
	str=str2 //可以赋值
	fmt.Printf("type=%T,str=%s\n",str,str)


	printLog(getFuncTwoAppend());
}
func getFuncTwoAppend()funcTwo{
	myAppend:=func(a,b int )string{
		return strconv.Itoa(a)+strconv.Itoa(b)
	}
	return myAppend
}
func printLog(theAppend funcTwo) {
	fmt.Println(theAppend(1,2));
}
type funcTwo func(int,int)string //方便函数做参数和返回值来传递