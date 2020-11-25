package main

import (
	"fmt"

	"github.com/maxdev/go-gingonic/entity"
	"github.com/maxdev/go-gingonic/repository"
)

func main() {

	repo := repository.NewRepo()

	newData := entity.Todo{
		Content: "data",
		Title:   "title",
	}

	newData2 := entity.Todo{
		Content: "data2",
		Title:   "title2",
	}

	res, err := repo.AddTodo(&newData)
	if err != nil {
		fmt.Print(err)
		return
	}
	res2, err2 := repo.AddTodo(&newData2)
	if err != nil {
		fmt.Print(err2)
		return
	}

	fmt.Println("res >", res)
	fmt.Println("res >", res2)
}
