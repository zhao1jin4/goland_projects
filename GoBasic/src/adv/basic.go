package main

import (
	"fmt"
)

func main() {
	cn:="中"
	man:=true;
	fmt.Printf("type=%T,cn=%q,a(ASCII)=%d man=%t\n",cn,cn,'a',man)
	//fmt.Printf("%c\n",cn)//%c有点复杂

	//fmt.Printf("请输入两个数\n")
	//var x,y int
	//fmt.Scan(&x,&y)
	//fmt.Printf("输入为x=%d,y=%d\n",x,y)


	//fmt.Printf("请输入一个整数，一个小数\n")
	//var w int
	//var h float32
	//fmt.Scan(&w,&h)
	//fmt.Printf("输入为w=%d,h=%f\n",w,h) //发现小数22.123精度不准


	//fmt.Printf("请输入一段字\n")
	//reader:=bufio.NewReader(os.Stdin)
	//s1,err:=reader.ReadString('\n')
	//if  err == nil {
	//	fmt.Printf("读到为=%s\n",s1) //
	//}

	 //同JS
	 func (a,b int ){
		 fmt.Println(a,b)
	 }(10,20)



	a1:=10
	defer calcSum(a1,20,30) //defer表示要所在函数main执行完成,在退出前再执行声明的函数calcSum,但参数已经传递了,只是晚点执行
	fmt.Printf("first\n")
	a1+=100;
	defer calcSum(a1,60) //如有多个defer，执行是是栈的方式，后进先出
	fmt.Printf("main method end\n")
}
func calcSum(nums ... int ) {//接收可变参数
	fmt.Printf("nums=%d\n",nums)
}