package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gobeam/stringy"

	"github.com/maxdev/go-gingonic/entity"
)

type todoRepositotyInDB struct {
	DB *sql.DB
}

func CreateRepositoryInDB(db *sql.DB) RepositoryI {
	return todoRepositotyInDB{DB: db}
}

func (tr todoRepositotyInDB) AddTodo(todo *entity.Todo) (int64, error) { // stmt, err := tr.DB.Prepare("S")
	stmt, err := tr.DB.Prepare("INSERT INTO todo (TITLE,CONTENT) VALUE (?,?)")

	if err != nil {
		return 0, errors.New("Can't Add todo stmt")
	}

	result, err := stmt.Exec(todo.Title, todo.Content)

	if err != nil {
		return 0, errors.New("Can't Add todo result")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("Can't Add todo id")
	}

	return id, nil
}

func (tr todoRepositotyInDB) GetTodos() ([]entity.Todo, error) {

	rows, err := tr.DB.Query("SELECT * FROM `todo`")
	defer rows.Close()
	if err != nil {
		return nil, errors.New("No data")
	}

	var response []entity.Todo

	for rows.Next() {
		var row entity.Todo

		err := rows.Scan(
			&row.Id,
			&row.Title,
			&row.Content,
			&row.IsDone,
			&row.CreateAt,
		)

		if err != nil {
			return nil, errors.New("No Data")
		}

		response = append(response, row)
	}

	return response, nil
}

func (tr todoRepositotyInDB) UpdateTodo(id int64, todo map[string]interface{}) (entity.Todo, error) {

	var emptyTodo entity.Todo

	attr := make([]string, 0, len(todo))

	for key, value := range todo {
		key := stringy.New(key)

		snakeValue := fmt.Sprint(key.SnakeCase("?", "").ToUpper())

		attr = append(attr, snakeValue+"="+fmt.Sprintf("'%v'", value))
	}

	attrJoin := fmt.Sprint(strings.Join(attr, ", "))

	sql := "UPDATE todo SET " + attrJoin + " WHERE todo.ID = ?"

	fmt.Println(sql)

	stmt, err := tr.DB.Prepare(sql)

	if err != nil {
		return emptyTodo, errors.New("SQL COMMAND ERROR")
	}

	_, errstmt := stmt.Exec(id)

	if errstmt != nil {
		return emptyTodo, errors.New("ERROR CAN'T EXCUTE COMMAND")
	}

	return emptyTodo, nil
}

func (tr todoRepositotyInDB) DeleteTodo(id int64) (string, error) {

	stmt, err := tr.DB.Prepare("DELETE FROM todo WHERE ID = ?;")

	if err != nil {
		return "", errors.New("ลบไม่ได้ไอเวร")
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return "", errors.New("ลบไม่ได้ไอเวร")
	}

	fmt.Println(result)

	return "ลบได้และ", nil
}

func (tr todoRepositotyInDB) GetByID(id int64) (entity.Todo, error) {

	sql := "SELECT * FROM todo WHERE ID = " + fmt.Sprint(id)

	row := tr.DB.QueryRow(sql)

	var todo entity.Todo

	err := row.Scan(
		&todo.Id,
		&todo.Title,
		&todo.Content,
		&todo.IsDone,
		&todo.CreateAt,
	)

	if err != nil {
		return entity.Todo{}, errors.New("Not found ID")
	}

	fmt.Println(todo)

	return todo, nil
}
