package v1

import "fmt"

// Init todo manager
func Init() *DataBase {
	db := newDatabase()
	err := db.open()
	if err != nil {
		fmt.Println(err)
	}

	db.autoMigrate()
	return db
}

// CreateTodo creates a todo with given title and completed flag.
func CreateTodo(database *DataBase, title string, completed bool) uint {
	todo := todoModel{
		Title:     title,
		Completed: completed,
	}

	database.save(&todo)
	return todo.ID
}
