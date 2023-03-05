package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB
var err error

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

	cmd := `insert into persons (name, age) values (?,?)`
	_, err := Db.Exec(cmd, "Tanaka", 49)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("insert success")

	//p := Person{}
	//cmd := `SELECT * FROM persons WHERE name = ?`
	//row := Db.QueryRow(cmd, "sano")
	//err := row.Scan(&p.Name, &p.Age)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		log.Println("No row")
	//	} else {
	//		log.Println(err)
	//	}
	//}
	//fmt.Println(p.Name, p.Age)

	//cmd := `select * from persons`
	//rows, _ := Db.Query(cmd)
	//defer rows.Close()
	//var pp []Person
	//
	//for rows.Next() {
	//	var p Person
	//	err := rows.Scan(&p.Name, &p.Age)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	pp = append(pp, p)
	//}
	//
	//for _, person := range pp {
	//	fmt.Println(person.Name, person.Age)
	//}
}
