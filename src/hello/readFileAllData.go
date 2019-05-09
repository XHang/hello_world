package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filePatch := os.Args[1:]
	if len(filePatch) == 0 {
		fmt.Print("无文件路径,请重新制定")
		return
	}
	counts := make(map[string]int)
	for _, patch := range filePatch {
		data, err := ioutil.ReadFile(patch)
		if err != nil {
			fmt.Printf("打开文件[%s]失败,原因是%v", patch, err)
			continue
		}
		lines := strings.Split(string(data), "\r\n")
		for _, line := range lines {
			counts[line]++
		}

		for line, count := range counts {
			fmt.Printf("打印出来的行是[%s],出现的次数是[%d] \n", line, count)
		}

	}

}
