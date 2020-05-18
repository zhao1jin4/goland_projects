package main

import (
	"fmt"
)

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

/*
Go 语言通过内置的错误接口
type error interface {
	Error() string
}*/
// 实现 `error` 接口,有错误时调用这个方法
//属于 DivideError结构体中的方法,可以是指针
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
	` //保留原来的换行符
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (int,  error) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,//结构初始化可指定key
			divider: varDivider,
		}
		return 0,&dData
	} else {
		return varDividee / varDivider, nil
	}
}

func main() {

	// 正常情况
	if result, err  := Divide(100, 10); err == nil {//err 变量范围是这个if块
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, err := Divide(100, 0); err != nil {
		fmt.Println("errorMsg is: ", err)
	}
	
	_, err:= Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		var err error  =fmt.Errorf("%f不能<0",f)
		//var err error =errors.New("math: square root of negative number")//errors是Go的返回一个error
		return 0,err
	}
	return f/2 , nil
}
