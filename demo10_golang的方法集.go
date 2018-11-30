package main

import (
	"fmt"
)

/*
	考点：golang的方法集
	解答：
	编译不通过！ 做错了！？说明你对golang的方法集还有一些疑问。
	一句话：golang的方法集仅仅影响接口实现和方法表达式转化，与通过实例或者指针调用方法无关。
*/

type people0 interface {
	Speak(string) string
}

type stduent0 struct{}

func (stu *stduent0) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo people0 = stduent0{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

/*
	解决方案：
	将方法中接受者类型修改为实例而非指针
	或
	在后续实例的时候，将实例变量定义为指针，以此达到方法接受者要求
*/
