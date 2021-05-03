//https://github.com/360EntSecGroup-Skylar/excelize
package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("D:/tmp/Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	cell, err := f.GetCellValue("Sheet1", "B2") //指定单元格
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
