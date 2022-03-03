package main

import (
	"fmt"
)

//事件列表
var event = make(map[string][]func(interface{}))

//注册
func register(name string, method func(interface{})) {
	list := event[name]
	list = append(list, method)
	event[name] = list
}

//调用
func call(name string, param interface{}) {
	list := event[name]
	for _,  method := range list {
		method(param)
	}
}

func do(i interface{}){
	fmt.Printf("method2: %v", i)
}

func main() {
	
	register("send", func(i interface{}){
		fmt.Println("method1: no param")
	})

	register("send", do)

	call("send", []int{1, 3, 5})
}