package main

import (
	"fmt"
	"sync"
)

/*
考点：map线程安全
解答：
可能会出现fatal error: concurrent map read and map write. 修改一下看看效果
*/

type UserAges struct {
	ages  map[string]int
	mutex sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.mutex.Lock()
	defer ua.mutex.Unlock()
	ua.ages[name] = age
}

/*
	//错误：未解开互斥锁
	func (ua *UserAges) Get(name string) int {
		if age, ok := ua.ages[name]; ok {
			return age
		}
		return -1
	}
*/

func (ua *UserAges) Get(name string) int {
	ua.mutex.Lock()
	defer ua.mutex.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {

	ua := UserAges{make(map[string]int), sync.Mutex{1, uint32(1)}}
	ua.Add("哈士奇", 6)
	ret := ua.Get("哈士奇")
	fmt.Println(ret)
}
