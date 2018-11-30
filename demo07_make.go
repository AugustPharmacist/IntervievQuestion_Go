package main

import "fmt"

/*
	考点：make默认值和append
	解答：
	make初始化是由默认值的哦，此处默认值为0
*/

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
