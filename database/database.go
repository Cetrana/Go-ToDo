package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"todoApi/model"
)

func getConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.WithError(err).Error("Cannot connect to database")
	}
	return db
}

func Todos() ([]model.Todo, error) {
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.WithError(err).Error("can't prepare statement")
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.WithError(err).Error("fail to get rows")
		return nil, err
	}
	var todos []model.Todo
	for rows.Next() {
		var id int
		var title, status string
		var todo model.Todo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Status)
		if err != nil {
			log.WithError(err).Error("can't scan row into var")
			return nil, err
		}
		todos = append(todos, todo)
		fmt.Println("one row", id, title, status)
	}
	return todos, nil
}

func Todo(queryId string) (model.Todo, error) {
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.WithError(err).Error("can't prepare statement")
		return model.Todo{}, err
	}

	row := stmt.QueryRow(queryId)
	var id int
	var title, status string

	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.WithError(err).Error("can't scan row into var")
		return model.Todo{}, err
	}

	return model.Todo{Id: id, Title: title, Status: status}, nil
}
func UpdateTodo(updateTodo model.Todo) error {
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE todos SET title=$2, status=$3 where id=$1")
	if err != nil {
		log.WithError(err).Error("can't prepare statement")
		return err
	}

	if _, err := stmt.Exec(updateTodo.Id, updateTodo.Title, updateTodo.Status); err != nil {
		log.WithError(err).Error("cannot update")
		return err
	}

	fmt.Println("update succes")
	return nil
}

func InsertTodo(todo model.Todo) model.Todo {
	db := getConnection()
	defer db.Close()

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1,$2) RETURNING id", todo.Title, todo.Status)
	var id int
	err := row.Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	todo.Id = id
	fmt.Println("Insert todos success id : ", id)
	return todo
}
func DeleteTodo(deleteID string) error {
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare("DELETE todos where id=$1")
	if err != nil {
		log.WithError(err).Error("can't prepare statement")
		return err
	}

	if _, err := stmt.Exec(deleteID); err != nil {
		log.WithError(err).Error("cannot delete")
		return err
	}

	fmt.Println("update succes")
	return nil
}
