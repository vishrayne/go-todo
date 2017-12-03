package v1

import (
	"errors"
	"fmt"
	"time"
)

//TodoManager representation
type (
	TodoManager struct {
		database *DataBase
	}

	// ResponseTodo - sanitized response format
	ResponseTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
		Created   time.Time
	}
)

// Init todo manager
func Init(shouldAutoMigrate bool, databasePath string) *TodoManager {
	db := newDatabase(databasePath)
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

// FindTodo updates a given todo
func (t *TodoManager) FindTodo(todoID uint) (ResponseTodo, error) {
	var todo todoModel
	t.database.findBy(&todo, "id = ?", todoID)

	if todo.ID == 0 {
		return ResponseTodo{}, errors.New("no record found")
	}

	return ResponseTodo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		Created:   todo.CreatedAt,
	}, nil
}

// UpdateTodo updates the given todo
func (t *TodoManager) UpdateTodo(todoID uint, title string, completed bool) (uint, error) {
	var todo todoModel
	t.database.findBy(&todo, "id = ?", todoID)

	if todo.ID == 0 {
		return 0, errors.New("no record found")
	}

	attrs := make(map[string]interface{})
	attrs["title"] = title
	attrs["completed"] = completed

	t.database.updateAttributes(&todo, attrs)
	return todoID, nil
}

// DeleteTodo deletes the given todo
func (t *TodoManager) DeleteTodo(todoID uint) error {
	var todo todoModel
	t.database.findBy(&todo, "id = ?", todoID)

	if todo.ID == 0 {
		return errors.New("no record found")
	}

	t.database.delete(&todo)
	return nil
}
