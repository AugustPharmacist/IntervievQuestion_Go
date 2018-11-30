package main

import "fmt"

/*
	解析
	考点：函数返回值命名
	在函数有多个返回值时，只要有一个返回值有指定命名，其他的也必须有命名。
	如果返回值有有多个返回值必须加上括号；
	如果只有一个返回值并且有命名也需要加上括号；
	此处函数第一个返回值有sum名称，第二个未命名，所以错误。
*/

func main() {
	sum, err := funcMui(2, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(sum)
}

func funcMui(x, y int) (sum int, err error) {
	return x + y, err
}

/*
	//错误代码
	func funcMui(x,y int)(sum int,error){
    	return x+y,nil
	}

*/
