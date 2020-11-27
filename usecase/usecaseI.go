package usecase

import "github.com/maxdev/go-gingonic/entity"

type UsecaseI interface {
	AddTodo(todo *entity.Todo) (int64, error)
	GetTodos() []entity.Todo
	UpdateTodo(id int64, todo *entity.Todo) (entity.Todo, error)
	DeleteTodo(id int64) (string, error)
}
