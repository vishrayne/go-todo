package todo

import (
	"testing"
)

func setup(t *testing.T) (*Manager, func(t *testing.T)) {
	t.Log("setting up...")
	todoManager := Init(TestMode, true)

	return todoManager, func(t *testing.T) {
		t.Log("tearing down...")
		// delete everything!
		todoManager.DeleteAll()
	}
}

func TestCreate(t *testing.T) {
	todoManager, tearDown := setup(t)
	defer tearDown(t)

	cases := []struct {
		name  string
		title string
		done  bool
	}{
		{"todo_1_test", "todo 1", true},
		{"todo_2_test", "todo 2", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			id := todoManager.Create(tc.title, tc.done)
			todo, _ := todoManager.Find(id)
			if todo.Title != tc.title {
				t.Errorf("expected todo title as %s but got todo => %s", tc.title, todo.Title)
			}

			if todo.Completed != tc.done {
				t.Errorf("expected todo completed as %t but got todo => %t", tc.done, todo.Completed)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	todoManager, tearDown := setup(t)
	defer tearDown(t)

	cases := []struct {
		name  string
		title string
		done  bool
	}{
		{"todo_1_test", "todo 1", true},
		{"todo_2_test", "todo 2", false},
		{"todo_3_test", "todo 3", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			todoManager.Create(tc.title, tc.done)
		})
	}

	todos := todoManager.GetAll()
	if len(todos) != len(cases) {
		t.Errorf("Expected %d todos to be created but was only %d", len(cases), len(todos))
	}
}

func TestUpdate(t *testing.T) {
	todoManager, tearDown := setup(t)
	defer tearDown(t)

	wantTitle := "updated title"
	wantCompleted := false

	id := todoManager.Create("title", true)
	todoManager.Update(id, wantTitle, wantCompleted)
	updatedTodo, _ := todoManager.Find(id)

	if updatedTodo.Title != wantTitle {
		t.Errorf("expected todo title as %s but got todo => %s", wantTitle, updatedTodo.Title)
	}

	if updatedTodo.Completed != wantCompleted {
		t.Errorf("expected todo completed as %t but got todo => %t", wantCompleted, updatedTodo.Completed)
	}
}
