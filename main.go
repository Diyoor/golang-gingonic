package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  http.StatusOK,
			"message": "Hello GIN",
		})
	})

	server.Run(":3001")
}
