package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Go是静态语言，编译时类型已经确定,动态类型要求是接口,每个接口会记录(pair)值(ValueOf)，和类型(ValueOf,即%T)

	var pi float64 =3.14 //任何类型都可以看到是一空接口类型，如float32也是
	fmt.Println("type=",reflect.TypeOf(pi)) //TypeOf参数如是空接口返回nil
	fmt.Println("value=",reflect.ValueOf(pi)) //ValueOf参数如是空接口返回0
	//接口变量 -> 反射对象
	var valReflect reflect.Value = reflect.ValueOf(pi)
	fmt.Println("is float32=",valReflect.Kind() == reflect.Float64)
	fmt.Println("Type=",valReflect.Type() )
	fmt.Println("Float=",valReflect.Float() )//如float32类型精度不准3.140000104904175，默认是float64

	// 反射对象 -> 接口变量
	convertVal:=valReflect.Interface().(float64)//对已知类型
	fmt.Println("convertVal=",convertVal)

	//--指针,如要修改值一定要使用指针
	var valPointerReflect reflect.Value = reflect.ValueOf(&pi)
	fmt.Println("valPointerReflect=",valPointerReflect)//是地址
	fmt.Println("valType=",valPointerReflect.Type()) //*float64

	//Elem修改值
	valElem:=valPointerReflect.Elem()//如不是指针这里报错
	fmt.Println("Elem=",valElem)//是值
	fmt.Println("CanSet=",valElem.CanSet())//是否可以修改值
	valElem.SetFloat(4.18) //修改值
	fmt.Println("after reflect set =",pi)

	//
	convertPointer:=valReflect.Interface().(float64) //不能是*float64
	fmt.Println("convertPointer=",convertPointer)//还是3.14 ???

}
