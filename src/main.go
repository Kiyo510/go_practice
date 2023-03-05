package main

import (
	"fmt"
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

	//err = user.CreateTodo("test content")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	t, err := models.GetTodo(5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(t)

	t.CONTENT = "updated!!!!"
	err = t.UpdateTodo()
	if err != nil {
		log.Fatalln(err)
	}

	todos, err := models.GetTodos()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(todos)

	//todos, _ = user.GetTodoByUser()
	//fmt.Println(todos)

}
