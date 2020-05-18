package main

import "fmt"

func main() {
	//----Map 无序的　是使用 hash 表来实现的
	var countryCapitalMap map[string]string  //map[key_data_type]value_data_type
	countryCapitalMap = make(map[string]string)

	countryCapitalMap [ "France" ] = "Paris"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	 for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	//查看元素在集合中是否存在
	captial, ok := countryCapitalMap [ "美国" ]
	if (ok) {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在")
	}
 	countryCapitalMap2 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	delete(countryCapitalMap2, "France")//自带的delete函数
	fmt.Println(countryCapitalMap2)
}


