package main

import (
	"fmt"
	"log"

	todo "github.com/vishrayne/go-todo"
)

func main() {
	todoManager := todo.Init(todo.DebugMode, true)

	id := todoManager.Create("title_1", true)
	//show all todo
	fmt.Println(todoManager.GetAll())

	activeTodo, err := todoManager.Find(id)
	if err != nil {
		log.Fatalf("unable to continue => %v", err)
	}

	fmt.Println(activeTodo)
	// update all fields
	fmt.Println(todoManager.Update(activeTodo.ID, "changed_title", false))
	fmt.Println(todoManager.Find(activeTodo.ID))
	// update only a single field
	fmt.Println(todoManager.Update(activeTodo.ID, "changed_title", true))
	fmt.Println(todoManager.Find(activeTodo.ID))
	// delete a todo
	todoManager.Delete(activeTodo.ID)
	fmt.Println(todoManager.GetAll())
}
