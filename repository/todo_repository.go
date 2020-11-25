package repository

import (
	"fmt"

	"github.com/maxdev/go-gingonic/entity"
)

type TodoRepoInterface struct {
	Todos []*entity.Todo
}

func NewRepo() TodoRepoInterface {
	return TodoRepoInterface{}
}

func (t *TodoRepoInterface) AddTodo(todo *entity.Todo) (id int64, err error) {

	if len(t.Todos) == 0 {
		todo.Id = 1
	} else {
		todo.Id = t.Todos[len(t.Todos)-1].Id + 1
	}

	t.Todos = append(t.Todos, todo)
	for _, value := range t.Todos {
		fmt.Print(value)
	}
	return todo.Id, nil
}
