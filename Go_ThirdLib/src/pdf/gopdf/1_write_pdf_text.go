//https://github.com/signintech/gopdf

package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	//var ttf = "c:\\WINDOWS\\Fonts\\SIMHEI.TTF" //设置了字体，中文正常
	var ttf = "./SIMHEI.TTF"
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("cnFont", ttf)
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("cnFont", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "您好")
	pdf.WritePdf("d:/tmp/hello.pdf")

}
