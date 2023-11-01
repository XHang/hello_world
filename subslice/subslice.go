package subslice

import (
	"fmt"
	"money-leak/util"
	"net/http"
	"runtime"
	"strconv"
)

var s0 []int

// G 有点像那个字串切割，这个数组切割也会导致内存泄露。one time produce about 22 MB
func leak() []int {
	// Assume the length of s1 is much larger than 30.
	s1 := util.Slice.CreateSliceWithRandomValues(10000)
	s0 = s1[len(s1)-30:]
	return s0
}
func Test() {
	m := make(map[int][]int)
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		times := request.URL.Query().Get("times")
		t, err := strconv.Atoi(times)
		if err != nil {
			fmt.Println("Error parsing times:", err)
			return
		}
		for i := 0; i < t; i++ {
			m[util.Random.GenerateNumber()] = Fix()
		}
		runtime.GC()
	})
	http.HandleFunc("/test1", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range m {
			fmt.Println(fmt.Println("k:", k, "v:", v))
		}
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}

// Fix occupy a few memory
func Fix() []int {
	s1 := util.Slice.CreateSliceWithRandomValues(10000)
	s0 = make([]int, 30)
	copy(s0, s1[len(s1)-30:])
	return s0
}
