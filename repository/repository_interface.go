package repository

import "github.com/maxdev/go-gingonic/entity"

type RepositoryInterface interface {
	AddTodo(todo *entity.Todo) (id int64, err error)
}
