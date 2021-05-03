package main

import (
	"fmt"
	"reflect"
)

func add(a, b int) int {
	return a + b
}
func main() {
	//-----普通函数,不是在struct里的
	funcValue := reflect.ValueOf(add)
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	retList := funcValue.Call(paramList)
	fmt.Println(retList[0].Int())
	//----------struct  tag
	type cat struct {
		Name string
		// 带有结构体tag的字段
		Type int `json:"type" id:"100"`
		//` 开始和结尾的字符串。这个字符串在Go语言中被称为 Tag（标签）。一般用于给字段添加自定义信息，方便其他模块根据信息进行不同功能的处理
		//键与值使用冒号分隔，值用双引号括起来；键值对之间使用一个空格分隔

	}

	ins := cat{Name: "mimi", Type: 1}
	typeOfCat := reflect.TypeOf(ins)
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
		if val, ok := catType.Tag.Lookup("id"); ok {
			fmt.Println("结构体标签id的值为", val)
		} else {
			fmt.Println("结构体标签不存在")
		}
	}
	//---IsNil  isValid
	// *int的空指针
	var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())
	s := struct{}{}
	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())
	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(s).MethodByName("").IsValid())
	// 实例化一个map
	m := map[int]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
	//---CanAddr
	{
		x := 2                   // value type variable?
		a := reflect.ValueOf(2)  // 2 int no
		b := reflect.ValueOf(x)  // 2 int no
		c := reflect.ValueOf(&x) // &x *int no
		d := c.Elem()            // 2 int yes (x)
		fmt.Println(a.CanAddr()) // "false"
		fmt.Println(b.CanAddr()) // "false"
		fmt.Println(c.CanAddr()) // "false"
		fmt.Println(d.CanAddr()) // "true"
		//Addr() 类似于语言层&操作
		//Elem()  类似于语言层*操作

	}
	//---TypeOf再New()返回Value
	{
		var a int
		// 取变量a的反射类型对象
		typeOfA := reflect.TypeOf(a)
		// 根据反射类型对象创建类型实例
		aIns := reflect.New(typeOfA)
		// 输出Value的类型和种类
		fmt.Println(aIns.Type(), aIns.Kind())
	}
}
