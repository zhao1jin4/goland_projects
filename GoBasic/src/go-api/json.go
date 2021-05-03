package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	//encoding/json 包性能比较低,ffjson (https://github.com/pquerna/ffjson) 性能高些,API使用是一样的
	array := [5]int{1, 2, 3, 4, 5}
	//array
	jsonByte, err := json.Marshal(array)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(string(jsonByte))
	//---map
	pair := make(map[string]float32)
	pair["age"] = 30
	jsonByte, err = json.Marshal(pair)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(string(jsonByte))
	//---struct 只有大写字母开头属性才会转换
	type Person struct {
		Name     string `json:"fullName"` //定义输出名字
		Birthday time.Time
		Weight   float32
	}
	lisi := Person{
		Name:     "lisi",
		Birthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Weight:   70.5,
	}
	jsonByte, err = json.Marshal(lisi)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(string(jsonByte)) //{"fullName":"lisi","Birthday":"2009-11-10T23:00:00Z","Weight":70.5}
	var emptyInter interface{}
	json.Unmarshal(jsonByte, &emptyInter)
	fmt.Printf("unmashal=%v \n", emptyInter) //结果是一个Map
	//unmashal=map[Birthday:2009-11-10T23:00:00Z Weight:70.5 fullName:lisi]

}
