package mygomock

import "fmt"

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	// ...
	var res = f.Bar(99)
	fmt.Println(res)


	res = f.Bar(101)
	fmt.Println(res)
}
