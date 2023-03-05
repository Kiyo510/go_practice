package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB
var err error

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:4306)/go_practice?charset=utf8mb4")
	if err != nil {
		fmt.Println("MySQL connection ERROR!!")
		fmt.Println(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Println("connetion refused!!")
		log.Fatal(err)
	} else {
		fmt.Println("connection success!!")
	}

	defer Db.Close()

	//cmd := `SELECT * FROM persons WHERE name = ?`
	//row := Db.QueryRow(cmd, "hiko")
	//err = row.Scan(&p.Name, &p.Age)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		log.Println("No row")
	//	} else {
	//		log.Println(err)
	//	}
	//}
	//fmt.Println(p.Name, p.Age)

	cmd := `select * from persons`
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	var pp []Person

	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Id, &p.Name, &p.Age)
		if err != nil {
			log.Fatal(err)
		}

		pp = append(pp, p)
	}

	for _, person := range pp {
		fmt.Println(person.Id, person.Name, person.Age)
	}
}
