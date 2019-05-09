//从标准输入中读取数据，输入到map数据类型中，并打印出来
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//声明一个空的map，key是字符串，value是int
	counts := make(map[string]int)
	//拿到一个输入对象
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入输入数组的元素，每一个元素输入文完毕，用回车键结束，结束输入，请输入eof，并回车")
	//其实
	for input.Scan() {
		str := input.Text()
		if str == "eof" {
			break
		}
		//每次输入时,输入的字符串,都作为key,value都自增
		counts[str]++
	}
	fmt.Println(counts)
	//遍历的顺序的乱序的
	for key, value := range counts {
		fmt.Printf("当前便利的map的ｋｅｙ是%s \n", key)
		fmt.Printf("当前便利ｍａｐ的ｖａｌｕｅ是%d", value)
		fmt.Println()
	}

}
