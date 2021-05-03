/*
生成二维码
https://github.com/skip2/go-qrcode

go get -u github.com/skip2/go-qrcode
*/
package main

import (
	"image/color"
	"io/ioutil"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	var png []byte
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256) //生成二进制数组

	err = ioutil.WriteFile("D:/tmp/qrcode.png", png, os.ModePerm)
	if err != nil {
		panic(err)
	}
	//直接写文件
	err = qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "D:/tmp/qrcode2.png")

	//黑色背景，白色为码
	// err = qrcode.WriteColorFile("https://example.org", qrcode.Medium, 256, color.Black, color.White, "D:/tmp/qrcode3.png")
	err = qrcode.WriteColorFile("https://example.org", qrcode.Medium, 256, color.RGBA{0, 255, 0, 200}, color.RGBA{0, 0, 255, 200}, "D:/tmp/qrcode3.png")
	if err != nil {
		panic(err)
	}
}
