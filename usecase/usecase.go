package usecase

import (
	"errors"

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
