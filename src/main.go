package main

import (
	"fmt"
	"go-practice/app/models"
	"log"
)

func main() {
	//err := controllers.StartMainServer()
	//if err != nil {
	//	log.Fatalln(err)
	//}

	user, err := models.GetUserByEmail("test3@example.com")
	if err != nil {
		log.Fatalln(err)
	}

	var session models.Session
	session, err = user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)

	result, err := session.CheckSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)
}
