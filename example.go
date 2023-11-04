package main

import (
	"money-leak/timeticker"
	_ "net/http/pprof"
)

// 实际在golang的处理中，我们并不怎么需要注意内存泄露问题
// 因为golang是自动垃圾回收的。
// 但实际上有几种场景是会导致内存泄露的。
// 这个项目的几个package将会介绍内存泄露的场景
func main() {
	timeticker.TimeTicker.TestFix()
}

// how to compare two functions:
//GO VERSION:go version go1.20 windows/amd64
//TEST TOOL:go tool pprof --inuse_space http://localhost:8080/debug/pprof/heap then type "web" to see the result.
