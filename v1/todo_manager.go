package v1

import (
	"fmt"
	"time"
)

//TodoManager representation
type TodoManager struct {
	database *DataBase
}

// ResponseTodo - sanitized response format
type ResponseTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Created   time.Time
}

// Init todo manager
func Init(shouldAutoMigrate bool) *TodoManager {
	db := newDatabase()
	err := db.open()
	if err != nil {
		fmt.Println(err)
	}

	if shouldAutoMigrate {
		db.autoMigrate()
	}

	return &TodoManager{
		database: db,
	}
}

// CreateTodo creates a todo with given title and completed flag.
func (t *TodoManager) CreateTodo(title string, completed bool) uint {
	todo := todoModel{
		Title:     title,
		Completed: completed,
	}

	t.database.save(&todo)
	return todo.ID
}

// GetAllTodo returns all stored todo records.
func (t *TodoManager) GetAllTodo() []ResponseTodo {
	var todos []todoModel
	t.database.find(&todos)

	if len(todos) <= 0 {
		return nil
	}

	var responseTodos []ResponseTodo
	for _, item := range todos {
		responseTodos = append(responseTodos, ResponseTodo{
			ID:        item.ID,
			Title:     item.Title,
			Completed: item.Completed,
		})
	}

	return responseTodos
}
