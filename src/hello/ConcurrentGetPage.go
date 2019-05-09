//并发获取页面信息
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	ch := make(chan string)
	//声明一个不定长的数组(切片)
	var urls []string
	fmt.Println("请输入url,输入eof结束")
	input := bufio.NewScanner(os.Stdin)
	for {
		input.Scan()
		url := input.Text()
		if url == "eof" {
			break
		}
		urls = append(urls, url)
		fmt.Println("成功录入一个骚东西")
	}
	fmt.Print("开始抓取页面")
	startTime := time.Now()
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("耗时秒数%f  \n", time.Since(startTime).Seconds())

}

//抓取网页,计算抓取时长
//第一个参数是字符串,第二个参数是channel,只可以写哦,如果声明成 ch<- chan string 则是可读
func fetch(url string, ch chan<- string) {
	fmt.Printf("开始抓取页面[%s] \n", url)
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("请求网页[%s]失败,因为[%v]", url, err)
		return
	}
	byteCount, err := io.Copy(ioutil.Discard, resp.Body)
	ch <- fmt.Sprintf("抓取网页[%s]成功,[%d]字节数已拿到,耗时[%f]", url, byteCount, time.Since(startTime).Seconds())

}
