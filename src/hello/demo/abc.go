package demo

import "fmt"

var abc string

func Abcs() {
	var a = []string{"1", "2", "3"}
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
}
func modify(arr []string) {
	arr[1] = "1111"
}
