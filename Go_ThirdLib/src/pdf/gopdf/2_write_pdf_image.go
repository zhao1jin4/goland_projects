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
	var err error
	err = pdf.AddTTFFont("loma", ttf)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Image("react.png", 200, 50, nil) //print image
	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetX(250) //move current location
	pdf.SetY(200)
	pdf.Cell(nil, "网页") //print text

	pdf.WritePdf("D:/tmp/image.pdf")

}
