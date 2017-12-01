package main

import (
	"fmt"
	"log"

	todo "github.com/vishrayne/go-todo/v1"
)

func main() {
	todoManager := todo.Init(true)

	id := todoManager.CreateTodo("title_1", true)
	//show all todo
	fmt.Println(todoManager.GetAllTodo())

	activeTodo, err := todoManager.FindTodo(id)
	if err != nil {
		log.Fatalf("unable to continue => %v", err)
	}

	fmt.Println(activeTodo)
	// update all fields
	fmt.Println(todoManager.UpdateTodo(activeTodo.ID, "changed_title", false))
	fmt.Println(todoManager.FindTodo(activeTodo.ID))
	// update only a single field
	fmt.Println(todoManager.UpdateTodo(activeTodo.ID, "changed_title", true))
	fmt.Println(todoManager.FindTodo(activeTodo.ID))
	// delete a todo
	todoManager.DeleteTodo(activeTodo.ID)
	fmt.Println(todoManager.GetAllTodo())
}
