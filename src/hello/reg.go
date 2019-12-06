package main

import (
	"fmt"
	"regexp"
)

func main() {
	//1 给定源字符串和正则表达式，判断是否匹配
	sourceStr := `1083594261@qq.com@ewretrey`
	matched, err := regexp.MatchString(`@\w+`, sourceStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", matched) // true

	reg, err := regexp.Compile(`@\w+`)
	if err != nil {
		panic(err)
	}
	isMatch := reg.MatchString(sourceStr)
	fmt.Println(isMatch)

	//打印出正则表达式
	fmt.Println(reg.String())

	//找到这个正则表达式第一个匹配的字符串
	fmt.Println(string(reg.Find([]byte(sourceStr))))

	fmt.Println(reg.FindAllStringSubmatch(sourceStr, 0))

	//匹配所有的字符串以及分组
	//返回的是一个嵌套数组的数组，[][]string
	//外层数组，包含的是每一个匹配的结果（一个字符串可能有多个子串都能匹配上正则表达式）
	//内层数组的第一个元素是整个匹配的结果，后面的元素是分组的结果
	//参数-1指表示查找所有匹配项，如果是非负数，则表示查找前n个匹配项
	reg = regexp.MustCompile(`(\w)(\w+)`)
	fmt.Printf("%q \n", reg.FindAllStringSubmatch("Hell Warld!", -1))
	//[["Hell" "H" "ell"] ["Warld" "W" "arld"]]

	//相当于FindAllStringSubmatch  n参数选择1，只查找第一个匹配项
	fmt.Println(reg.FindStringSubmatch("Hell Warld!"))

	//查找所有匹配项，不包括子匹配项，也就是分组的结果，n参数同上
	fmt.Println(reg.FindAllString("Hell Warld!", 1))

	//查找第一个匹配项，不包括子匹配项，也就是分组的结果，不匹配返回空，如果匹配的本身是空字符串，且匹配成功，也是空字符串，对于这两种情况，调用者需要区分
	fmt.Println(reg.FindString("Hell Warld!"))

}
