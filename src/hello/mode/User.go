package mode

import "fmt"

type User struct {
	UserName string
	Password string
	ID       int
	Sex      string
}

func (u *User) SetUserName(name string) {
	u.UserName = name
}
func ShowUserName(user *User) {
	fmt.Println(user.UserName)
}
