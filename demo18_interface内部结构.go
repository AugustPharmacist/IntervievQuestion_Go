package main

import "fmt"

/*
	解析
	考点：interface内部结构

	non-empty interface
*/

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	var x *int = nil
	Foo(x)
}
