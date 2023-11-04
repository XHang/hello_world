package timeticker

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

type timeTicker struct{}

var TimeTicker timeTicker

// Example1 time.NewTicker可以创建一个定时器，当它被创建时，每隔指定的几秒内，都会往里面的channel塞一个时间
// 这个定时器就可以被其他的协程所使用，从中取出时间，然后做一些事情
// 但是，如果要释放这个定时器里面的资源，必须调用它的Stop方法，否则会造成定时器的内存泄漏
func (t *timeTicker) Example1() {
	ticker := time.NewTicker(time.Second * 3)
	go func() {
		for {
			select {
			case ca := <-ticker.C:
				fmt.Printf("tick... %v \r\n", ca)
			}
		}
	}()
	for {
		time.Sleep(time.Second * 1)
	}
}
func (t *timeTicker) TestMemoryLeaking() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		for i := 0; i < 10000; i++ {
			a := time.NewTicker(time.Second * 3)
			fmt.Printf("it created %v", a)
		}
		runtime.GC()
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
func (t *timeTicker) TestFix() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		for i := 0; i < 10000; i++ {
			a := time.NewTicker(time.Second * 3)
			fmt.Printf("it created %v", a)
			a.Stop()
		}
		runtime.GC()
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
