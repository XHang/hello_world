package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	//每次循环，range返回两个数值，第一个是索引，第二个是索引的值。但是我们不需要索引，所以可以用_这个空白标识符来接受
	//另外，应该知道的是，go语言不支持定义变量，却不使用该变量的情况
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("接下来是见证奇迹的时刻，分开打印数组和角标")
	loop()
	fmt.Print("继续见证奇迹，使用join函数来打印数组")
	join_exampple()
}
func loop() {
	for index, value := range os.Args[0:] {
		fmt.Print("当前索引的角标")
		fmt.Print(index)
		fmt.Print("当前索引的值")
		fmt.Print(value)
		fmt.Println()
	}
}
func join_exampple() {
	fmt.Println(strings.Join(os.Args, " "))
}
