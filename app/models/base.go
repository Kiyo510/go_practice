package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go-practice/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var Db *sql.DB
var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, "root:root@tcp(127.0.0.1:4306)/"+config.Config.DbName+"?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
    id INT(11) AUTO_INCREMENT NOT NULL,
    uuid VARCHAR(255) NOT NULL ,
    name VARCHAR(255) NOT NULL ,
    email VARCHAR(255) NOT NULL ,
    password VARCHAR(255) NOT NULL ,
    created_at datetime default current_timestamp,
    PRIMARY KEY (id))`, tableNameUser)

	_, err := Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
    id INT(11) AUTO_INCREMENT NOT NULL,
    content TEXT NOT NULL ,
    user_id INT(11) NOT NULL ,
    created_at datetime default current_timestamp,
    PRIMARY KEY (id))`, tableNameTodo)

	_, err = Db.Exec(cmdT)
	if err != nil {
		log.Fatalln(err)
	}
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
