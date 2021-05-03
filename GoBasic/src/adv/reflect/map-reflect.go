package main

import (
	"fmt"
	"reflect"
)

func main() {
	//----map反射
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "c": 3, "b": 2}
	fmt.Println(reflect.DeepEqual(m1, m2))

	iter := reflect.ValueOf(m1).MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		fmt.Println("key=", k, "value=", v)
	}

	mapVal := reflect.ValueOf(m1)
	if mapVal.MapIndex(reflect.ValueOf("a")).IsValid() { //SetMapIndex
		mapVal.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(11)) //SetMapIndex
		fmt.Println(m1)
	}

}
