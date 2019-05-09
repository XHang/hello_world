package main

import (
	"hello_world/src/hello/demo"
	"html/template"
	"log"
	"os"
)

const templ = `{{.Username}} friend:
{{range .Friend}}----------------------------------------
UaseName: {{.Username}}
ID: {{.ID}}
Password: {{.Password}}
Age:{{.Age}}
NextAge:{{.Age|nextAge}}
{{end}}`

var report = template.Must(template.New("User").
	Funcs(template.FuncMap{"nextAge": nextAge}).
	Parse(templ))

func main() {
	var user = demo.User{
		ID:       12,
		Username: "张三",
		Password: "李四",
		Age:      123,
	}
	var friend = user
	user.Friend[0] = &friend
	if err := report.Execute(os.Stdout, user); err != nil {
		log.Fatal(err)
	}
}
func nextAge(age int32) (int32, int) {
	return age + 1, 1
}
