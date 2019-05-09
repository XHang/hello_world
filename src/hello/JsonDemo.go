package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Id    int
	Title string
	//结构体的json序列化和反序列化时,参照的字段是released
	Year int `json:"released"`
	//如果Color字段为零值,比如说false,则序列化成json时不会给出这个字段
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	movies := initMovie()
	//没有格式化的json数据
	//fmt.Print(toJsonString(movies))
	//经过格式化的数据
	jsonStr := toJsonFormatString(movies)
	fmt.Println("结构体转字符串")
	fmt.Println(jsonStr)
	fmt.Println("字符串转结构体")
	fmt.Println(jsonToObject(jsonStr))
}

func initMovie() []Movie {
	m := []Movie{
		{Id: 1, Title: "复联4", Year: 2019, Color: false, Actors: []string{"你", "我"}},
		{Id: 2, Title: "菊花侠大战菊花怪", Year: 189, Color: false, Actors: []string{"夏明", "东海"}},
	}
	return m
}
func toJsonString(m []Movie) string {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("对象转json失败,原因是[%v]", err)
	}
	return string(data)
}
func toJsonFormatString(m []Movie) string {
	data, err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		fmt.Printf("对象转json失败,原因是[%v]", err)
	}
	return string(data)
}

//json字符串转结构体
func jsonToObject(jsonStr string) []Movie {
	var movie []Movie
	//第二个参数必须是指针,否则会报non-pointer
	if err := json.Unmarshal([]byte(jsonStr), &movie); err != nil {
		fmt.Printf("字符串转结构体失败,因为[%v]", err)
	}
	return movie
}
