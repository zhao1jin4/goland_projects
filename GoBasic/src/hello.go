package main //main函数所在文件的包名必须是main,同一文件夹的go文件包名必须全一样,建议用文件夹名

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")

	//如一行多写条语句用;分隔
	var a = "中文"
	var b string = "english" //类型前没有:
	var c bool
	println(a, b, c) //变量被使用

	var myint = 10
	myint2 := 10 //前面不能是已经定义的变量

	println(myint, myint2)

	//--data type
	var myint8 int8 = 10
	var myint16 int16 = 10
	var myint32 int32 = 10
	var myint64 int64 = 10
	println(myint8, myint16, myint32, myint64)

	var myuint8 uint8 = 10
	var myuint16 uint16 = 10
	var myuint32 uint32 = 10
	var myuint64 uint64 = 10
	println(myuint8, myuint16, myuint32, myuint64)

	var myfloat32 float32 = 10
	var myfloat64 float64 = 10
	println(myfloat32, myfloat64) //+1.000000e+001 ,也有计算不准的情况

	var mybyte byte = 0x20 //类似 uint8
	var myrune rune = 10   //类似 int32
	var myuint uint = 10   //32 或 64 位
	var myint3 int = 10    //与 uint 一样大小

	println(mybyte, myrune, myuint, myint3)

	//var myptr uintptr=&myuint;
	//println(myptr)

	var x, y = 20, 20
	println(x, y)

	var x1, y1 uint16 = 30, 30
	println(x1, y1)

	a1, b1 := 4, 5 //只能出现在函数中
	println(a1, b1)

	const a2, b2, c2 = 2, 4, "22" //常量

	var //一般用于声明全局变量
	(
		width  uint32
		height uint32
	)
	println(width, height)
	//string 属于值类型

	//常量还可以用作枚举：
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	const (
		myio   = iota //0
		myio2         //1
		myio3         //2
		myioR  = iota //3
		myioR1        //4

	)
	fmt.Println(myio, myio2, myio3, myioR, myioR1)

	const (
		i = 1 << iota
		j = 3 << iota
		k // 3<<2
		l // 3<<3。
	)
	fmt.Println("i=", i) //1
	fmt.Println("j=", j) //6
	fmt.Println("k=", k) //12
	fmt.Println("l=", l) //24
	//&取址，*指针，有struct 同C ,有goto,只有for循环

	//switch
	var marks int = 90
	var grade string = "B"
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	var xIterface interface{}
	switch i := xIterface.(type) { //强转为type，表示可用switch判断类型，后也可用i继续调用子类
	case nil: //匹配nil
		fmt.Printf(" x 的类型 :%T", i) //%T类型
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
	//使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false") //这个会被执行
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
	//select

	//---for
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}
	var size int = 5
	var iter int = 1
	for iter < size {
		iter++
		fmt.Printf("iter 的值为: %d\n", iter)
	}
	numbers := [6]int{1, 2, 3, 5} //数组
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
	//----function
	fmt.Println(max(200, 120)) //都是自己的函数
	strF, strT := swap("Mahesh", "Kumar")
	fmt.Println(strF, strT)

	var num1 int = 100
	var num2 int = 200
	fmt.Printf("交换前，num1 的值 : %d\n", num1)
	fmt.Printf("交换前，num2 的值 : %d\n", num2)
	swapPtr(&num1, &num2)
	fmt.Printf("交换后，num1 的值 : %d\n", num1)
	fmt.Printf("交换后，num2 的值 : %d\n", num2)

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	//闭包
	nextNumber := getSequence()
	fmt.Println(nextNumber()) //1
	fmt.Println(nextNumber()) //2
	fmt.Println(nextNumber()) //3

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1()) //1
	fmt.Println(nextNumber1()) //2

	//----array
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)
	var emptyArray [10]int
	var ii int
	for ii = 0; ii < 10; ii++ {
		emptyArray[ii] = ii + 100
	}

	var myArray = [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11}}
	fmt.Println(myArray)

	//----struct
	fmt.Println(Books{"Go 语言", "li", "Go 语言教程", 6495407})
	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "li"})

	var book1 Books //使用类型总是在后面

	book1.title = "Go 语言"
	book1.author = "li"
	book1.subject = "Go 语言教程"
	book1.book_id = 6495407
	printBook(book1)
	printBookPtr(&book1)

	var struct_pointer *Books
	struct_pointer = &book1
	fmt.Println(struct_pointer.author) //不同于C 不是->

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea()) //结构体中的方法

	//---切片Slice("动态数组")
	var numbers2 = make([]int, 3, 5) //make函数创建 len,cap
	printSlice(numbers2)

	slice2 := []int{1, 2, 3} //和数组的不同是  []中为空，即没有...和数字
	printSlice(slice2)

	var myIntArray = [...]int{10, 20, 30}
	sliceFromArray := myIntArray[0:3] //[startIndex:endIndex]
	//sliceFromArray := myIntArray[0:]
	//sliceFromArray := myIntArray[:3]
	printSlice(sliceFromArray)

	var empSlice []int
	if empSlice == nil { //ObjectiveC
		fmt.Printf("切片是空的")
	}
	fmt.Println("sliceFromArray[1:3] ==", sliceFromArray[1:3])   //20,30 从索引(包含) 到索引(不包含)
	var sliceFromArrayRes []int = append(sliceFromArray, 50, 60) //增加无素
	fmt.Println("sliceFromArrayRes =", sliceFromArrayRes)

	newSlice := make([]int, len(sliceFromArrayRes), (cap(sliceFromArrayRes))*2) //len,cap函数
	copy(newSlice, sliceFromArrayRes)                                           //拷贝  dest,src
	printSlice(newSlice)
	//----
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(15))) //类型转换
}

//闭包 返回一个函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}
type Circle struct {
	radius float64
}

//属于 Circle结构体中叫方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
func max(num1, num2 int) int { //func 如{在新的一行就报错
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//函数返回多个值
func swap(x, y string) (string, string) {
	return y, x
}
func swapPtr(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

//func getAverage(arr []int, size int) float32 //函数参数数组，像C

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
func printBookPtr(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func Factorial(n uint64) (result uint64) { //返回参数名
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
