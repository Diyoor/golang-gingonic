package repository

import (
	"github.com/maxdev/go-gingonic/entity"
)

type RepositoryI interface {
	AddTodo(todo *entity.Todo) (int64, error)
}
