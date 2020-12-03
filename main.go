package main

import (
	"strconv"

	"github.com/maxdev/go-gingonic/pkg/database"

	"github.com/maxdev/go-gingonic/dto"
	"github.com/maxdev/go-gingonic/entity"
	"github.com/maxdev/go-gingonic/repository"
	"github.com/maxdev/go-gingonic/usecase"

	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	repo := repository.CreateRepositoryInDB(database.DB)
	uc := usecase.CreateTodoUsecase(repo)

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {

		res, err := uc.GetTodos()

		if err != nil {
			c.JSON(400, gin.H{
				"errors": err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"data": res,
		})
	})

	server.POST("/", func(c *gin.Context) {

		req := dto.AddTodo{}
		c.ShouldBindJSON(&req)
		data := entity.Todo{
			Content:  req.Content,
			Title:    req.Title,
			CreateAt: time.Now(),
		}
		res, err := uc.AddTodo(&data)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"res": res,
		})

	})

	server.PATCH("/:id", func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		var reqUpd map[string]interface{}

		c.ShouldBindJSON(&reqUpd)

		res, err := uc.UpdateTodo(id, reqUpd)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"res": res,
		})

	})

	server.DELETE("/:id", func(c *gin.Context) {

		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		res, err := uc.DeleteTodo(id)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{
			"res": res,
		})

	})

	server.Run(":3001")
}
