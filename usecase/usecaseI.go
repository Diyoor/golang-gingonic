package usecase

import "github.com/maxdev/go-gingonic/entity"

type UsecaseI interface {
	AddTodo(todo *entity.Todo) (int64, error)
}
