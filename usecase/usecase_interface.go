package usecase

import "github.com/maxdev/go-gingonic/entity"

type UsecaseInterface interface {
	AddTodo(todo *entity.Todo) (id int64, err error)
}
