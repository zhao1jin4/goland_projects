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

	//pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})//简单的

	//找开要密码，可以禁止复制
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
		Protection: gopdf.PDFProtectionConfig{
			UseProtection: true,
			Permissions:   gopdf.PermissionsPrint | gopdf.PermissionsCopy | gopdf.PermissionsModify,
			OwnerPass:     []byte("123456"),
			UserPass:      []byte("123456789")},
	})

	pdf.AddPage()
	err := pdf.AddTTFFont("times", ttf)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Image("react.png", 200, 50, nil)
	err = pdf.SetFont("times", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetX(30)
	pdf.SetY(40)
	pdf.Text("Link to example.com")
	pdf.AddExternalLink("http://example.com/", 27.5, 28, 125, 15)

	pdf.SetX(30)
	pdf.SetY(70)
	pdf.Text("到第二页")
	pdf.AddInternalLink("anchor", 27.5, 58, 120, 15)

	pdf.AddPage()
	pdf.SetX(30)
	pdf.SetY(100)
	pdf.SetAnchor("anchor")
	pdf.Text("Anchor position")

	//画线
	pdf.SetLineWidth(2)
	pdf.SetLineType("dashed")
	pdf.Line(10, 30, 585, 30)
	//画椭圆
	pdf.SetLineWidth(1)
	pdf.Oval(100, 200, 500, 500)
	//画多边形
	pdf.SetStrokeColor(255, 0, 0)
	pdf.SetLineWidth(2)
	pdf.SetFillColor(0, 255, 0)
	pdf.Polygon([]gopdf.Point{{X: 10, Y: 30}, {X: 585, Y: 200}, {X: 585, Y: 250}}, "DF")

	//旋转文字，如做水层式半透明背景水印？
	pdf.SetX(100)
	pdf.SetY(100)
	pdf.Rotate(270.0, 100.0, 100.0)
	pdf.Text("X公司...")
	pdf.RotateReset() //不旋转，没用？

	/* 半透明，报错？
	transparency := gopdf.Transparency{
		Alpha:         0.5,
		BlendModeType: "",
	}
	pdf.SetTransparency(transparency)
	*/
	pdf.WritePdf("d:/tmp/hello.all.pdf")

}
