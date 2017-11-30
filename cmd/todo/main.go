package main

import (
	"fmt"

	todo "github.com/vishrayne/go-todo/v1"
)

func main() {
	todoManager := todo.Init(true)
	todoManager.CreateTodo("title_2", true)
	fmt.Println(todoManager.GetAllTodo())
}
