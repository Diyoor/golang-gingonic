package usecase

import (
	"errors"
	"fmt"

	"github.com/maxdev/go-gingonic/entity"
	"github.com/maxdev/go-gingonic/repository"
)

type TodoUsecase struct {
	repo repository.RepositoryI
}

func CreateTodoUsecase(repo repository.RepositoryI) UsecaseI {
	return &TodoUsecase{repo: repo}
}

func (uc *TodoUsecase) AddTodo(todo *entity.Todo) (int64, error) {

	id, err := uc.repo.AddTodo(todo)

	if err != nil {
		return 0, errors.New("Invalid Data !")
	}

	return id, nil
}

func (uc *TodoUsecase) GetTodos() []entity.Todo {
	data := uc.repo.GetTodos()
	return data
}
func (uc *TodoUsecase) UpdateTodo(id int64, todo *entity.Todo) (entity.Todo, error) {

	// var req entity.Todo

	// oldData := uc.repo.GetTodos()

	// for _, data := range oldData {
	// 	if data.Id == id {
	// 		req = data

	// 		fmt.Println(data)
	// 	}
	// }

	fmt.Println(id)
	fmt.Println(todo)

	data, err := uc.repo.UpdateTodo(id, todo)

	return data, err
}

func (uc *TodoUsecase) DeleteTodo(id int64) (string, error) {
	data, err := uc.repo.DeleteTodo(id)
	return data, err
}
