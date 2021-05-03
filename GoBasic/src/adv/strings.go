package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.ToLower("Gopher"))
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))

	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) //%q安全去除单/双引号
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
}
