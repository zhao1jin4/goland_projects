//https://github.com/360EntSecGroup-Skylar/excelize

package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet2")
	f.SetCellValue("Sheet2", "A2", "Hello world.") //按sheet名字去定位，不太好，而不是对象
	f.SetCellValue("Sheet1", "B2", 100)
	//如何写下拉外表的选项
	f.SetActiveSheet(index)
	if err := f.SaveAs("d:/tmp/Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
