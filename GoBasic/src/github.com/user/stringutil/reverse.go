//包名和目录名建议相同,这里不同
package stringutil1

import "fmt"

func init(){ //init函数在导入包，GO语言自动调用，做初始化工作，和main参数一样，不能有参数和返回值
	//每个包中都可以有init,可以重复出现,顺序是从上到下，同一文件夹按文件名排序，不同包按导入顺序
	//如main依赖A ->B -> C ,执行顺序为C->B->A->main
 	fmt.Println("stringutil1包Init函数初始了")
}
func init(){
	fmt.Println("stringutil1包Init函数初始2")
}
//大写开头 表示可被其它包导入使用，小写开头表示，只为自己包使用
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

type Persion struct {
	Name string
	age int //小写开头包外不可见
}
