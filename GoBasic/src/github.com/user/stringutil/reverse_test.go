package stringutil1

import (
	"fmt"
	"testing"
)
func init(){
	fmt.Println("stringutil1包Init函数初始Test")// xx_test.go 不会被main调用
}
//另一种函数类型为 func BenchmarkXxx(b* testing.B)
func TestReverse(t *testing.T) { //函数名TestXxx，参数t *testing.T
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
