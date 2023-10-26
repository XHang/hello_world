package main

import (
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

// 实际在golang的处理中，我们并不怎么需要注意内存泄露问题
// 因为golang是自动垃圾回收的。
// 但实际上有几种场景是会导致内存泄露的。
// 后面的几个函数将显示这种场景
func main() {
	arr := make([]string, 0)
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		for i := 0; i < 1000; i++ {
			arr = append(arr, FixFunctionOfLack2(RandStringRunes(1<<20)))
		}
		runtime.GC()
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}

func RandStringRunes(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	characters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	arr := make([]rune, n)
	for i := 0; i < n; i++ {
		arr[i] = characters[r.Intn(len(characters))]
	}
	return string(arr)
}

// FunctionOfLack1 当进行字串操纵时，生成后的字串会和原先的字符串共享同一个内存块，这是golang的小技巧
// 但是这样也造成了一个问题，如果你不再需要原字符串了。即使你把原字符串归解除引用。实际的原字符串还是没有消失
// 因为他们共享同一个内存块
func FunctionOfLack1(src string) string {
	s := src[0:50]
	return s
}

// FixFunctionOfLack2 为了解决这个问题。首先可以先把得到字串转为字节，然后再转字符串。从根本上解决掉分享内存的弊端
func FixFunctionOfLack2(src string) string {
	return string([]byte(src[0:50]))
}

//以下是结论。
//通过两个函数的对比后发现，每执行一次。FunctionOfLack1就回增加将近1GB的内存。而FixFunctionOfLack2则是一直在500KB左右徘徊。可见内存泄露确实是存在的
//GO VERSION:go version go1.20 windows/amd64
//TEST TOOL:go tool pprof --inuse_space http://localhost:8080/debug/pprof/heap then type "web" to see the result.
