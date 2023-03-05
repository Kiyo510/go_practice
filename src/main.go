package main

import (
	"go-practice/app/models"
	"log"
)

func main() {
	//fmt.Println(u)
	//
	//err := u.CreateUser()
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//user, err := u.GetUser(1)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(user)
	//
	//user.Name = "test3"
	//user.Email = "test3@example.com"
	//err = user.UpdateUser()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	var user models.User
	var err error
	user, err = user.GetUser(1)
	if err != nil {
		log.Fatalln(err)
	}

	t := models.Todo{
		CONTENT: "todo test",
		UserID:  user.ID,
	}

	err = t.CreateTodo()
	if err != nil {
		log.Fatalln(err)
	}

	t, err = t.GetTodo(1)
}
