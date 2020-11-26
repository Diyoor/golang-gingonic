package usecase

import (
	"github.com/maxdev/go-gingonic/entity"
	"github.com/maxdev/go-gingonic/repository"
)

type Usecase struct {
	repo repository.TodoRepoInterface
}

func NewUsecase(repo repository.TodoRepoInterface) UsecaseInterface {
	return Usecase{repo: repo}
}

func (uc Usecase) AddTodo(todo *entity.Todo) (id int64, err error) {
	uc.repo.AddTodo(todo)
	return
}
