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
	if todos, err := database.Todos(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, todos)
	}

}

func getTodo(c *gin.Context) {
	if todo, err := database.Todo(c.Param("id")); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func postTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.JSON(http.StatusOK, database.InsertTodo(todo))

}

func putTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	if err := database.UpdateTodo(todo); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

func deleteTodo(c *gin.Context) {
	if err := database.DeleteTodo(c.Param("id")) ; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	}

}
