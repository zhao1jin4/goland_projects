package main

import "fmt"

type Human struct {
	name string
	age int
	weight float32
}
type Student struct {
	Human //没有名字，就是名字和类一样，同JS,也可以成为提升字段
	grade string
	age float32 //同名，类型变了
}

type Move interface{ //接口只能定义方法声明(也可没有方法，像java的Serialize)，可以做为方法参数
	walk()
}
type Move2 interface{
	walk()//不同接口可有同名方法，哪些有实现呢
}
type Advance interface{
	Move //类似struct来继承，可以加多个
	work()
}

func printInterface(m Move ){
	if s,ok:=m.(*Student);ok { //断言，强转为子类
		s.walk()
	}else {
		m.walk();
	}
	//方试二 强转为子类
	switch s:=m.(type) {
		case  *Student:
	  		s.walk()
		default:
			s.walk();
	}
}
func noNameInterface(m interface{}){ //匿名空接口
	fmt.Println(m) //fmt.Println的参数也是空接口
}
func (h  Human )walk(){
	fmt.Println(h.name +"is is human , walk")
}
func (h  Student )walk(){
	fmt.Println(h.name +"is is Student , walk")
}

func (h  Human )work(){
	h.age=33
	fmt.Println(h.name +"is is human , working")
}
func (h * Student )work(){ //属于不同的结构体的方法可以同名，这里是重写父类方法，指针传递的是地址,里面修改会影响外面
	h.age=33.5
	fmt.Println(h.name +"is is Student , working")
}

func main(){
	p:=Human{name:"张三",age:20}
	fmt.Println(p.name,p.age)

	s:=Student{Human:Human{name:"李同学",age:18},grade:"A"}

	p.work()
	s.work()
	fmt.Println(p.name,p.age)
	fmt.Println(s.Human.name,s.grade)
	fmt.Println(s.name,s.age,s.grade)//可以直接访问name，模拟了继承

	s.walk();//子类可以调用父类的方法

	printInterface(&s)//如声明是指针，传递要为地址
	printInterface(s);
	printInterface(p);

	var  m Move =p //可以指向子类,模拟多态
    m.walk()

	s.work()
	var s1 *Student  =&s
	s1.work()

	noNameInterface(123);
	noNameInterface("abc");

	map1:=make(map[string]interface{})
	map1["name"]="李四"
	map1["age"]=20 //值可以是任何类型

	slice1:=make([]interface{},0,10) //切换数组任意类型
	slice1=append(slice1,"李四",30)

}
