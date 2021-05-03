package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
type RefInterface interface{
	say(msg string)
	printInfo()
}
*/
type Persion struct {
	Name string
	Age  int
}

func (p Persion) Say(msg string) { //可以是 (p Persion) 也可是 (p *Persion)
	fmt.Println("say:", msg)
}
func (p Persion) PrintInfo() { //方法名要大写
	fmt.Printf("name=%s,age=%d\n", p.Name, p.Age)
}

func printFunc() {
	fmt.Println("printFunc Invoked")
}
func calcFunc(a int, b int) string {
	res := a + b
	var resStr = strconv.Itoa(res)
	fmt.Printf("calc %d+%d=%d", a, b, res)
	return resStr
}

func main() {
	var p1 Persion = Persion{"李四", 28}
	//showReflect(p1) //有错误
	var p2 *Persion = &p1
	fmt.Printf("p2.Name=%s,(*p2).Name=%s\n", p2.Name, (*p2).Name) //都可以

	changeFieldByReflect(p2) //struct指针修改值
	//----无参数函数调用
	reflectFunc := reflect.ValueOf(printFunc)
	fmt.Println("printFunc Kind=", reflectFunc.Kind()) //func，必须是func才能调用Call
	fmt.Println("printFunc Type=", reflectFunc.Type()) //func()

	emptySlice := make([]reflect.Value, 0)
	resSlice := reflectFunc.Call(emptySlice) //调用方法，空参数传 nil 或 空切片
	fmt.Println("resSlice=", resSlice)

	//----有参数函数调用
	reflectFunc = reflect.ValueOf(calcFunc)
	fmt.Println("reflectFunc .Kind=  ", reflectFunc.Kind()) //func
	fmt.Println("reflectFunc Type= ", reflectFunc.Type())   // func(int, int) string

	paramSlice := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	resSlice = reflectFunc.Call(paramSlice) //调用方法，空参数传 nil 或 空切片
	fmt.Println("resSlice=", resSlice)
	res := resSlice[0].Interface().(string) //Interface函数  反射类型对象 -> 接口类型变量 ，返回空接口类型强转为string
	fmt.Println("res=", res)

	///---对象的方法
	//refType:=reflect.TypeOf(p2)
	refValue := reflect.ValueOf(p2) //指针
	fmt.Printf("refValue =%v \n", refValue)
	//refElem :=refValue.Elem()//指针的指针才要这一步
	valMethod := refValue.MethodByName("PrintInfo") //方法名一定要大写
	fmt.Printf("valMethod =%v \n", valMethod)       //%v显示值内存地址
	fmt.Printf("valMethod Kind=%s,Type=%s \n", valMethod.Kind(), valMethod.Type())

	emptySlice = make([]reflect.Value, 0)
	valMethod.Call(emptySlice) //调用方法，空参数传 nil 或 空切片

	valMethod = refValue.MethodByName("Say")
	paramSlice1 := []reflect.Value{reflect.ValueOf("参数1	")}
	valMethod.Call(paramSlice1)

}
func showReflect(common interface{}) {
	refType := reflect.TypeOf(common)
	fmt.Println("Name=", refType.Name()) //Persion
	fmt.Println("Kind=", refType.Kind()) //struct还有slice ,map,

	refValue := reflect.ValueOf(common)
	fmt.Println("Type=", refValue.Type()) //main.Persion
	fmt.Println("refValue=", refValue)    //{李四 28}
	fmt.Printf("---Field\n")
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		val := refValue.Field(i).Interface()                                                        //Interface()得到值
		fmt.Printf("index=%d,type=%s,name=%s,value=%v\n", field.Index, field.Type, field.Name, val) //%v显示通用类型的值
	}

	ageField, ok := refType.FieldByName("Age") //Type根据字段名取
	if !ok {
		fmt.Printf("没有Age字段\n")
	} else {
		fmt.Printf("Age type =%s\n", ageField.Type)
	}

	ageField2 := refValue.FieldByName("Age") //Value根据字段名取
	fmt.Printf("Age =%d\n", ageField2)

	fmt.Printf("---Type Method\n")
	for i := 0; i < refType.NumMethod(); i++ {
		method := refType.Method(i)
		fmt.Printf("TypeOf Name=%s,Type=%s V=%v \n", method.Name, method.Type, method.Type)
	}
	//---ValueOf得到不到Method??
	fmt.Printf("---Value Method\n")
	for i := 0; i < refValue.NumMethod(); i++ {
		method := refValue.Method(i)
		fmt.Printf("ValueOf Type=%s  V=%v \n", method.Type, method.Type) //没有 method.Name,显示不对  ???
	}
	valMethod := refValue.MethodByName("printInfo") //返回一个Value,错误???
	//valMethod:=refValue.MethodByName("say");//返回一个Value
	fmt.Printf("valMethod =%v \n", valMethod)
	//fmt.Printf("valMethod Kind=%s,Type=%s \n",valMethod.Kind(),valMethod.Type())

}
func changeFieldByReflect(p2 *Persion) {
	refType := reflect.TypeOf(p2)
	refValue := reflect.ValueOf(p2)
	fmt.Println("Type=", refValue.Type())      //*main.Persion
	fmt.Println("val Kind =", refValue.Kind()) //ptr
	fmt.Println("typ Kind =", refType.Kind())  //ptr
	if refValue.Kind() == reflect.Ptr {
		valElem := refValue.Elem()
		if valElem.CanSet() {
			fldAge := valElem.FieldByName("Age")
			fldAge.SetInt(30) //结构体的属性大写字母开头的才可写
			fmt.Println("after change age =", p2.Age)
		}
	}
}
