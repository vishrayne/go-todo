package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	todo "github.com/vishrayne/go-todo"
)

const todoManagerKey string = "todo_manager_key"

func main() {
	// gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(todoMiddleware())

	engine.GET("/", rootHandler)
	engine.GET("/ping", pingHandler)

	v1 := engine.Group("/api/v1/todos")
	{
		v1.GET("/", showAllTodoHandler)
		v1.POST("/create", createTodoHandler)
		v1.GET("/:id", showTodoHandler)
		v1.PUT("/:id", updateTodoHandler)
		v1.DELETE("/:id", deleteTodoHandler)
	}

	engine.Run(":8080")
}

func todoMiddleware() gin.HandlerFunc {
	// one-time initialization
	todoManager := todo.Init(todo.DebugMode, true)

	return func(c *gin.Context) {
		c.Set(todoManagerKey, todoManager)
		c.Next()
	}
}

func rootHandler(c *gin.Context) {
	pingHandler(c)
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Service is live!", "title": "todo-manager"})
}

func showAllTodoHandler(c *gin.Context) {
	todoManager := c.MustGet(todoManagerKey).(*todo.Manager)
	c.JSON(http.StatusOK, gin.H{"data": todoManager.GetAll()})
}

// TODO: cleanup
func createTodoHandler(c *gin.Context) {
	todoManager := c.MustGet(todoManagerKey).(*todo.Manager)
	title := c.PostForm("title")
	done := c.PostForm("done")

	completed, err := strconv.ParseBool(done)
	if err != nil {
		log.Printf("parsing bool failed, setting todo as incomplete -> %v", err)
		completed = false
	}

	id := todoManager.Create(title, completed)
	c.JSON(http.StatusCreated, gin.H{"message": "todo created", "id": id})
}

// TODO: cleanup
func showTodoHandler(c *gin.Context) {
	todoManager := c.MustGet(todoManagerKey).(*todo.Manager)
	id := c.Param("id")
	todoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("error parsing todo id -> %v", err)
		c.JSON(http.StatusNotFound, gin.H{"data": "", "error": err.Error()})
		return
	}

	activeTodo, err := todoManager.Find(uint(todoID))
	if err != nil {
		log.Printf("error fetching todo with id[%d] -> %v", todoID, err)
		c.JSON(http.StatusNotFound, gin.H{"data": "", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": activeTodo})
}

// TODO: cleanup
func updateTodoHandler(c *gin.Context) {
	todoManager := c.MustGet(todoManagerKey).(*todo.Manager)
	id := c.Param("id")
	title := c.PostForm("title")
	done := c.PostForm("done")

	todoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("error parsing todo id -> %v", err)
		c.JSON(http.StatusOK, gin.H{"data": "", "error": err.Error()})
		return
	}

	completed, err := strconv.ParseBool(done)
	if err != nil {
		log.Printf("parsing bool failed, not skipping -> %v", err)
		c.JSON(http.StatusOK, gin.H{"data": "", "error": "missing field `done`"})
		return
	}

	_, err = todoManager.Update(uint(todoID), title, completed)
	if err != nil {
		log.Printf("updating todo[%d] failed -> %v", todoID, err)
		c.JSON(http.StatusOK, gin.H{"data": "", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Updated successfully"})
}

// TODO: cleanup
func deleteTodoHandler(c *gin.Context) {
	todoManager := c.MustGet(todoManagerKey).(*todo.Manager)
	id := c.Param("id")

	todoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Printf("error parsing todo id -> %v", err)
		c.JSON(http.StatusOK, gin.H{"data": "", "error": err.Error()})
		return
	}

	err = todoManager.Delete(uint(todoID))
	if err != nil {
		log.Printf("error removing todo[%d] -> %v", todoID, err)
		c.JSON(http.StatusOK, gin.H{"data": "", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Deleted successfully"})
}
