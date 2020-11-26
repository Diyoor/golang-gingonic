package main

import (
	"fmt"

	"github.com/maxdev/go-gingonic/usecase"

	"github.com/maxdev/go-gingonic/entity"
	"github.com/maxdev/go-gingonic/repository"
)

func main() {

	repo := repository.NewRepo()
	uc := usecase.NewUsecase(repo)

	newData := entity.Todo{
		Content: "data",
		Title:   "title",
	}

	res, err := uc.AddTodo(&newData)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("res >", res)
}
