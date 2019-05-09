//从命令行接受一段数组，并打印出来
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, str string
	for i := 0; i < len(os.Args); i++ {
		s += str + os.Args[i]
		str = ""
	}
	fmt.Print(s)
}
