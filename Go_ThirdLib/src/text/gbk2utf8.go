package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {
	var fromFile string = "D:/tmp/my_gbk.txt"
	from, err := os.OpenFile(fromFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer from.Close()
	reader := bufio.NewReader(from)

	for {

		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		var mylen int = len(data) * 4
		decodeBytes := make([]byte, mylen, mylen)
		decodeBytes, err = simplifiedchinese.GB18030.NewDecoder().Bytes(data)
		if err != nil {
			break
		}
		var str string
		str = string(decodeBytes)
		fmt.Printf("转换为GBK后为=%s", str)
	}
}
