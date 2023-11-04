package finalizer

import (
	"fmt"
	"net/http"
	"runtime"
)

type finalizer struct{}

var Finalizer finalizer

type Human struct {
	Baby  *Human
	Block [1 << 20]int //once it created,it will occupy much memory
}

// Test 这是真正的内存泄露
// 通过调用它。 变量占用的内存不能回收。
// 因为这些变量相互引用
// 当然我们其实很少调用过SetFinalizer，这个东西一般是变量被回收的时候才会被调用。而且它还不能保证一定会执行
func (f *finalizer) Test() {
	mother := new(Human)
	baby := new(Human)

	mother.Baby = baby
	baby.Baby = mother

	var _ = func(t *Human) {
		fmt.Printf("%v \r\n", t)
	}
	//runtime.SetFinalizer(mother, finalizer)

}
func (f *finalizer) Run() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		for i := 0; i < 10; i++ {
			f.Test()
		}
		fmt.Printf("it runs \r\n")
		runtime.GC()
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
