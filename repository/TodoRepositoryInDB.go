package repository

import (
	"database/sql"
	"errors"
	"fmt"

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

func (tr todoRepositotyInDB) UpdateTodo(id int64, todo *entity.Todo) (entity.Todo, error) {
	var emptyTodo entity.Todo

	stmt, err := tr.DB.Prepare("UPDATE todo SET CONTENT  = ?, TITLE = ? ,IS_DONE = ? WHERE `todo`.`ID` = ?")
	if err != nil {
		return emptyTodo, errors.New("E")
	}
	result, err := stmt.Exec(id)

	if err != nil {
		return emptyTodo, errors.New("A")
	}

	fmt.Println(result)

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
