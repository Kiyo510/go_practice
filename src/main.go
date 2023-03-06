package main

import (
	"go-practice/app/controllers"
	"log"
)

func main() {
	err := controllers.StartMainServer()
	if err != nil {
		log.Fatalln(err)
	}
}
