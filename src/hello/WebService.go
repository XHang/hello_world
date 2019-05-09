//提供对外访问接口的服务器
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var counts int
var mutex sync.Mutex

func main() {
	//每次请求的时候,其实就是起一个线程,访问对应的处理函数,在这里,线程也被称为一个goroutine
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", count)
	http.HandleFunc("/getMsg", printMsg)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(write, "您访问的地址是[%s]", request.URL.Path)
}
func count(write http.ResponseWriter, request *http.Request) {
	//对线程中间共享变量进行管控
	mutex.Lock()
	counts++
	mutex.Unlock()
	fmt.Fprintf(write, "您访问了总共[%d]次", counts)
}

//将请求头和表单数据传给响应
func printMsg(write http.ResponseWriter, request *http.Request) {
	for key, value := range request.Header {
		fmt.Fprintf(write, "[%s]=[%s] \n", key, value)
	}
	fmt.Fprintln(write)
	//go语言专属的if语句,分号前面可以执行一个初始化语句
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range request.Form {
		fmt.Fprintf(write, "[%s]=[%s] \n", k, v)
	}
	fmt.Fprintf(write, "您访问的主机是[%s]  \n", request.Host)
	fmt.Fprintf(write, "您访问的地址是[%s]  \n", request.RemoteAddr)
}
