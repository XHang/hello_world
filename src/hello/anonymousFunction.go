package main

import "fmt"

func main() {
	m := map[string]itype{"add": new(Add), "delete": new(Delete)}
	funcs := make([]func(), 0)
	for _, i := range m {
		funcs = append(funcs, func() { fmt.Println(i.GetType()) })
	}
	for _, v := range funcs {
		v() //预期delete ，add  实际delete，因为匿名函数体使用的外部变量引用的是变量的地址，在循环中，这个变量的地址不断值新的值。
		//在最后的调用时，就用的是最后指向的值，如果希望匿名函数体用的变量是在调用时就固定的话，引用的变量必须是值，不能是指针类型

	}
}

type itype interface {
	GetType() string
}
type Add struct {
}

func (*Add) GetType() string {
	return "add"
}

type Delete struct {
}

func (*Delete) GetType() string {
	return "delete"
}
