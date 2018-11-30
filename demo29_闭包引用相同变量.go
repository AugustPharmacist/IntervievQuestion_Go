package main

/*
	解析
	考点：闭包引用相同变量*
	结果：

	100
	110
*/

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func main() {
	a, b := test(100)
	a()
	b()
}
