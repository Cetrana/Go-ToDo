package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoApi/database"
	"todoApi/model"
)

func StartServer(url string) {
	router := gin.Default()
	router.GET("/todos/:id", getTodo)
	router.GET("/todos", getTodos)
	router.PUT("/todos/:id", putTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.POST("/todos", postTodo)
	router.Run(url)
}

func getTodos(c *gin.Context) {
	c.JSON(200, database.Todos())
}

func getTodo(c *gin.Context) {
	c.JSON(200, database.Todo(c.Param("id")))
}

func postTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(200, database.InsertTodo(todo))

}

func putTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	if database.UpdateTodo(todo) {
		c.JSON(200, gin.H{"status": "success"})
	} else {
		c.JSON(400, gin.H{"status": "failed"})
	}
}

func deleteTodo(c *gin.Context) {
	if database.DeleteTodo(c.Param("id")) {
		c.JSON(200, gin.H{"status": "deleted"})
	} else {
		c.JSON(400, gin.H{"status": "failed"})
	}

}
