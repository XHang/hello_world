package demo

import "fmt"

func demo() {
	var arr = []string{"1", "2", "3"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, value := range arr {
		fmt.Println(value)
	}
}
