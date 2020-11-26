package main

import (
	"github.com/maxdev/go-gingonic/dto"
	"github.com/maxdev/go-gingonic/entity"
	"github.com/maxdev/go-gingonic/repository"
	"github.com/maxdev/go-gingonic/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	repo := repository.CreateRepository()
	uc := usecase.CreateTodoUsecase(repo)

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {

		var newData dto.AddTodo

		if err := c.ShouldBindJSON(&newData); err != nil {
			c.JSON(200, gin.H{
				"erroro": err,
			})
			return
		}

		saveTodo := entity.Todo{
			Content: newData.Content,
			Title:   newData.Title,
		}

		res, _ := uc.AddTodo(&saveTodo)

		c.JSON(200, gin.H{
			"data": newData,
			"id":   res,
		})
	})

	server.Run(":3001")
}
