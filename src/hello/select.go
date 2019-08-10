package main

import (
	"fmt"
	"time"
)

func main() {
	//select基本用法
	chan1 :=make(chan int)
	chan2 :=make(chan int)
	go func() {
		chan1<-1
	}()
	/*go func() {
		time.Sleep(4*time.Second)
		chan2<-2
	}()*/
	time.Sleep(1000)
	select {
	case chan1<- 1:
		fmt.Println("chan1接受到数据了")
	case <-chan2 :
		fmt.Println("chan2接受到数据了")
	default:
		// 如果成功向chan2写入数据，则进行该case处理语句
	}

}
