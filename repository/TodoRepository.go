package repository

import (
	"errors"

	"github.com/maxdev/go-gingonic/entity"
)

type TodoRepository struct {
	TodoList  []*entity.Todo // map[int64][]byte
	CurrentID int64
}

func CreateRepository() RepositoryI {
	return &TodoRepository{}
}

func (t *TodoRepository) AddTodo(todo *entity.Todo) (int64, error) {

	todo.Id = t.CurrentID + 1
	t.TodoList = append(t.TodoList, todo)
	t.CurrentID++

	return todo.Id, nil
}

func (t *TodoRepository) GetTodos() []entity.Todo {
	var todoLst []entity.Todo
	for _, todo := range t.TodoList {
		todoLst = append(todoLst, *todo)
	}
	return todoLst

}
func (t *TodoRepository) UpdateTodo(id int64, todo *entity.Todo) (entity.Todo, error) {
	var todoMod entity.Todo
	for i, _ := range t.TodoList {
		if int64(i) == id {
			t.TodoList[id-1] = todo
		}
	}

	todoMod = *todo

	return todoMod, errors.New("error")
}

func (t *TodoRepository) DeleteTodo(id int64) (string, error) {

	return "", nil
}
