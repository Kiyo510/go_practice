package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	CONTENT   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos(
                  content,
                  user_id,
                  created_at) values (?,?,?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `SELECT id,content,user_id,created_at FROM todos WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(&todo.ID, &todo.CONTENT, &todo.UserID, &todo.CreatedAt)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select * from todos`
	rows, err := Db.Query(cmd)
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.CONTENT, &t.UserID, &t.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, t)
	}

	return todos, err
}

func (u *User) GetTodoByUser() (todos []Todo, err error) {
	cmd := `select * from todos where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.CONTENT, &t.UserID, &t.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, t)
	}

	return todos, err
}

func (t Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, user_id = ? where id = ?`
	_, err := Db.Exec(cmd, t.CONTENT, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
