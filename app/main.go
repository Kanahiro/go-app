package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin/infra"
	"gin/service"
)

func main() {
	engine := gin.Default()
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	db := infra.InitDb()

	rootGroup := engine.Group("/api")
	{
		todoGroup := rootGroup.Group("/todo")
		{
			todoService := service.NewTodoService(*db)

			todoGroup.GET("/", func(c *gin.Context) {
				todoList := todoService.GetTodoList()
				c.JSON(http.StatusOK, todoList)
			})
			todoGroup.POST("/", func(c *gin.Context) {
				text := "sample"
				done := false
				todo := todoService.CreateTodo(text, done)
				c.JSON(http.StatusOK, todo)
			})
			todoGroup.PUT("/:id", func(c *gin.Context) {
				id, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					panic(err)
				}
				todo := todoService.UpdateTodoByID(id, "sample", false)
				c.JSON(http.StatusOK, todo)
			})
			todoGroup.DELETE("/:id", func(c *gin.Context) {
				id, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					panic(err)
				}
				result := todoService.DeleteTodoByID(id)
				c.JSON(http.StatusOK, result)
			})
		}
	}

	engine.Run(":3000")

}
