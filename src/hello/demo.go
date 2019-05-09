package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserName string
	Password string
	ID       int
	Sex      string
}

func main() {

	var u User
	ref := reflect.TypeOf(u)
	fieldCount := ref.NumField()
	for i := 0; i < fieldCount; i++ {
		field := ref.Field(i)
		fmt.Println("字段名：" + field.Name)
		//field还可以获取字段的tag，相当于字段的注解了
	}
}
