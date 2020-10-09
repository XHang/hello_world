package main

import "fmt"

//实验rune的各种特性
func main() {
	str := "我是你爷爷"
	fmt.Printf("中文字符串[%v]调用len查询长度，结果为[%v]\r\n", str, len(str))
	fmt.Printf("中文字符串[%v]转rune数组查询长度，结果为[%v]\r\n", str, len([]rune(str)))
	fmt.Printf("中文字符串[%v]取第一个字节，结果为[%v]\r\n", str, string(str[0]))
	fmt.Printf("中文字符串[%v]取第一个字符，结果为[%v]\r\n", str, string([]rune(str)[0]))

}
