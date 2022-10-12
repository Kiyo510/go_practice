package main

import "fmt"

type User struct {
	Name string
	Age  int
	//X,Y int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u *User) SetName(name string) {
	u.Name = name
}
func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	user1.SetName("hogetarou")
	fmt.Println(user1)

	user2 := NewUser("mike", 100)
	fmt.Println(user2)
}
