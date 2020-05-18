package main

import (
	"fmt"
	"strconv"
)
func main() {
	//fmt.Println("a"+3)//error
	var num,err=strconv.Atoi("20")
	if err == nil{
		fmt.Println(num*3)
	}
	s := strconv.Itoa(-42)
	fmt.Println(s+"123")


	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)
	fmt.Println(b,f,i,u)

	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-42, 16)//16进制
	s4 := strconv.FormatUint(42, 16)
	fmt.Println(s1,s2,s3,s4)
}
