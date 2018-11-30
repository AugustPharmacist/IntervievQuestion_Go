package main

/*
	考点：type

	编译失败，因为type只能使用在interface
*/

func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}

func GetValue() int {
	return 1
}

/*
	解决方案：
	定义接口，并在其中分别实现能接收到int，string，interface等返回值的方法
	type showInterface interface {
		xxx() int
		xxx() string
		xxx() interface
	}
*/
