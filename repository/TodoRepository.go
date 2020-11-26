package repository

import (
	"github.com/maxdev/go-gingonic/entity"
)

type TodoRepository struct {
	TodoList []*entity.Todo
}

func CreateRepository() RepositoryI {
	return &TodoRepository{}
}

func (t *TodoRepository) AddTodo(todo *entity.Todo) (int64, error) {

	var localData []*entity.Todo

	if len(localData) == 0 {
		todo.Id = 1
	} else {
		todo.Id = todo.Id + 1
	}

	localData = append(localData, todo)

	t.TodoList = localData

	return todo.Id, nil
}
