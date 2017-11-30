package main

import (
	todo "github.com/vishrayne/go-todo/v1"
)

func main() {
	db := todo.Init()
	todo.CreateTodo(db, "title_1", false)
}
