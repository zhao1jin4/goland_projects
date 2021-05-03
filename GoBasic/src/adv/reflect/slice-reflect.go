package main

import (
	"fmt"
	"reflect"
)

func main() {
	//---slice 反射
	var intSlice = []int{256, 512, 1024}
	intSliceElemValue := reflect.ValueOf(&intSlice).Elem()
	if intSliceElemValue.CanSet() {
		newSliceValue := []int{2560, 5120, 10240}
		newVale := reflect.ValueOf(newSliceValue)
		intSliceElemValue.Set(newVale)
		fmt.Println("NewSliceVal =", intSlice)
	}
	{
		var intSlice = []int{256, 512, 1024}
		intSliceValue := reflect.ValueOf(intSlice)
		e := intSliceValue.Index(0) //Index函数
		if e.CanSet() {
			e.SetInt(2560)
			fmt.Println("NewVal =", intSliceValue)
		}
	}

}
