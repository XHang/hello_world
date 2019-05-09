//演示结构体
package main

import "fmt"

//雇员结构体
//结构体内部成员不能包含自身,虽然它不保错...报的是运行时错
type Employee struct {
	Id       int
	Username string
	Password string
	Sex      string
	Address  string
	Company  string

	bossAddress *Employee
}

var my Employee

func main() {
	my := initEmployee()
	fmt.Println(toString(my))

}
func toString(e Employee) string {
	var sb = fmt.Sprintf("Id=[%d],Username=[%s],Password=[%s],sex=[%s],Address=[%s],Company=[%s] ", e.Id, e.Username, e.Password, e.Sex, e.Address, e.Company)
	if e.bossAddress != nil {
		sb += "boss={" + toString(*(e.bossAddress)) + "}"
	}
	return sb

}

//初始化结构体数据
func initEmployee() Employee {
	my.Address = "奋发园501"
	my.Company = "阿米巴网络科技有效公司"
	my.Id = 1
	my.Sex = "保密"
	//取结构体变量的指针(地址)
	position := &my.Password
	//把字符串存进这个地址
	*position = "1008611"
	//实际上,上面两个语句和my.Password="108611"是等价

	//以下两个语句是花样式赋值,首先是取地址,然后是取地址的存放位置赋给一个变量,其实就是兜了一圈而已
	var variable *Employee = &my
	variable.Password = "11111"
	c := my

	variable.bossAddress = &c
	return my
}
func Method(a int, b int) (string, int) {
	return fmt.Sprintf("%d%d", a, b), a + b
}
