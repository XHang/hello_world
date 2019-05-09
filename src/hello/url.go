//读取到一个网页,并将其打印出来
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		fmt.Println("请输入至少一个地址")
		os.Exit(1)
	}
	for _, url := range urls {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			fmt.Printf("该网站[%s]没有协议头  \n", url)
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("打开地址[%s]失败,原因是[%v]   \n", url, err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		fmt.Println("成功获取响应,响应的状态码为:" + resp.Status)
		//读完流要关闭啊亲
		resp.Body.Close()
		if err != nil {
			fmt.Printf("抽取该网页[%s]的输出流数据失败,原因是[%v]  \n", url, err)
		}
		fmt.Println("接下来展示改网页的内容啦")
		fmt.Println(string(b))
		fmt.Println("然后是自己定义的方法抽取响应")
		echoPage(url)

	}
}

//该函数也可以获取页面,参数是URL地址,返回值是字符串
func echoPage(url string) {
	if url == "" {
		fmt.Println("请输入至少一个地址")
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("打开地址[%s]失败,原因是[%v]   \n", url, err)
		return
	}
	//b,err:=ioutil.ReadAll(resp.Body)
	//将来自页面的读取流,拷贝到标准输出流中,直接打印在控制台,就不需要定义一个字节缓存区
	io.Copy(os.Stdout, resp.Body)
	//读完流要关闭啊亲
	resp.Body.Close()

}
