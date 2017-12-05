package todo

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	toml "github.com/pelletier/go-toml"
)

const (
	DebugMode   string = "dev"
	ReleaseMode string = "prod"
	TestMode    string = "test"
)

type (
	//Manager representation
	Manager struct {
		mode     string
		database *DataBase
		config   *toml.Tree
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
func Init(appMode string, shouldAutoMigrate bool) *Manager {
	mode := DebugMode
	switch appMode {
	case DebugMode:
	case ReleaseMode:
	case TestMode:
		mode = appMode
	default:
		panic("app mode unknown: " + appMode)
	}

	path, _ := filepath.Abs(".")
	absPath := filepath.Join(path, "config", "settings.toml")
	config, _ := toml.LoadFile(absPath)

	databaseConfig := config.Get("database").(*toml.Tree)
	databaseEnv := databaseConfig.Get(mode).(*toml.Tree)
	databasePath := filepath.Join(path, "data", databaseEnv.Get("name").(string))

	db := newDatabase(databasePath)
	err := db.open()
	if err != nil {
		fmt.Println(err)
	}

	if shouldAutoMigrate {
		db.autoMigrate()
	}

	return &Manager{
		database: db,
		mode:     mode,
		config:   config,
	}
}

// Create creates a todo with given title and completed flag.
func (t *Manager) Create(title string, completed bool) uint {
	todo := todoModel{
		Title:     title,
		Completed: completed,
	}

	t.database.save(&todo)
	return todo.ID
}

// GetAll returns all stored todo records.
func (t *Manager) GetAll() []ResponseTodo {
	var todos []todoModel
	t.database.find(&todos)

	if len(todos) <= 0 {
		return []ResponseTodo{}
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

// Find updates a given todo
func (t *Manager) Find(todoID uint) (ResponseTodo, error) {
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

// Update updates the given todo
func (t *Manager) Update(todoID uint, title string, completed bool) (uint, error) {
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

// Delete deletes the given todo
func (t *Manager) Delete(todoID uint) error {
	var todo todoModel
	t.database.findBy(&todo, "id = ?", todoID)

	if todo.ID == 0 {
		return errors.New("no record found")
	}

	t.database.delete(&todo)
	return nil
}

// DeleteAll deletes everything
func (t *Manager) DeleteAll() {
	t.database.deleteAll(&todoModel{})
}
