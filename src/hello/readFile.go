//从命令行参数中读取一个参数,并将该文件读到内存中,按每行打印出文件的内容
//几个坑
//1. 有些文件的文件编码,会导致文件含有bom头,于是便利出来的文件含有一些骚东西
//2. input.Scan()方法读取的是utf-8编码的文件,除此之外的文件编码读取会出更多的骚东西
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePaths := os.Args[1:]
	count := make(map[string]int)
	if len(filePaths) == 0 {
		fmt.Print("你输入的文件路径长度为0,请使用控制台录入文件信息")
		countLine(os.Stdin, count)
	} else {
		for _, path := range filePaths {
			fmt.Printf("开始打印该文件[%s]内容 \n", path)
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("文件[%s]打开失败,原因是%v", path, err)
				continue
			}
			fmt.Println("成功打开该文件")
			countLine(file, count)
			file.Close()
		}
	}

	for line, count := range count {
		if count < 1 {
			fmt.Println("有一个计数不存在的行出现啦")
			continue
		}
		fmt.Printf("拿到某行[%s],该行出现的次数是%d", line, count)
		fmt.Println()
	}

}
func countLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		msg := input.Text()
		fmt.Println("读取一行")
		//fmt.Printf("读取到[%s] \n",msg)
		if msg == "eof" {
			fmt.Println("读到末尾了,退出程序")
			break
		}
		counts[msg]++
	}
}
