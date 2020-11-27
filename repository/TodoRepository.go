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

func (t *TodoRepository) GetTodos() ([]entity.Todo, error) {
	var todoLst []entity.Todo
	for _, todo := range t.TodoList {
		todoLst = append(todoLst, *todo)
	}

	if len(todoLst) == 0 {
		return nil, errors.New("No data !")
	}

	return todoLst, nil

}
func (t *TodoRepository) UpdateTodo(id int64, todo *entity.Todo) (entity.Todo, error) {
	var data entity.Todo
	if id > int64(len(t.TodoList)) {
		return entity.Todo{}, errors.New("Out of length")
	} else {

		for i, v := range t.TodoList {
			if v.Id == id {
				var dataStruct entity.Todo

				dataStruct = *todo

				dataStruct.Id = v.Id
				dataStruct.CreateAt = v.CreateAt

				data = dataStruct
				t.TodoList[i] = &dataStruct

			}
		}

		return data, nil
	}
}

func (t *TodoRepository) DeleteTodo(id int64) (string, error) {

	return "", nil
}
