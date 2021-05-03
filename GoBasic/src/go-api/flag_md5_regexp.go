package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"regexp"
)

func main() {
	env := flag.String("env", "dev", "环境") //命令行启动，读参数
	pageSize := flag.Int("size", 30, "页大小")
	flag.Parse()
	fmt.Println("env=", *env, "size=", *pageSize)
	// go run src/arg_md5_regexp.go --help 显示参数是一个-,但如果参数过多，其实要用两个-
	//go run src/arg_md5_regexp.go -env=test -size=10
	//go run src/arg_md5_regexp.go --env=test --size=10
	//---------md5
	myMd5 := md5.New()
	myMd5.Write([]byte("内容"))
	result := myMd5.Sum([]byte(""))
	fmt.Printf("result=%x \n", result)

	//---正则 官方文档很多示例
	//reg := regexp.MustCompile(`(1{1}[0-9]{10}){1}`)
	reg := regexp.MustCompile(`(1[0-9]{10})+`)
	matched := reg.FindAllStringSubmatch("my phone 13012345678 in book,your is 18011112222", -1)
	fmt.Printf("matched=%v \n", matched) //%v	the value in a default format

	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))

}
