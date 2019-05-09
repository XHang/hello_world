package demo

type User struct {
	ID       int
	Username string
	Password string
	Age      int32
	Friend   [1]*User
}
