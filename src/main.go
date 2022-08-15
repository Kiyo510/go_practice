package main

import "fmt"

type User struct {
	name  string
	age   int
	email string
}

func NewPost() *User {
	u := new(User)
	u.name = "taro"
	u.age = 39
	return u
}
func main() {
	user := NewPost()
	fmt.Println(user)
}
