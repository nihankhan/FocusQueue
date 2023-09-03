package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoStruct(t *testing.T) {
	// Create a Todo instance
	todo := Todo{
		Id:        1,
		Item:      "Buy groceries",
		Completed: 0,
	}

	// Check the fields of the Todo struct
	assert.Equal(t, 1, todo.Id)
	assert.Equal(t, "Buy groceries", todo.Item)
	assert.Equal(t, 0, todo.Completed)
}

func TestViewStruct(t *testing.T) {
	// Create a list of Todo instances
	todos := []Todo{
		{Id: 1, Item: "Task 1", Completed: 0},
		{Id: 2, Item: "Task 2", Completed: 1},
		{Id: 3, Item: "Task 3", Completed: 0},
	}

	// Create a View instance
	view := View{
		Todos: todos,
	}

	// Check the fields of the View struct
	assert.Len(t, view.Todos, 3)
	assert.Equal(t, "Task 1", view.Todos[0].Item)
	assert.Equal(t, 1, view.Todos[1].Completed)
	assert.Equal(t, 3, view.Todos[2].Id)
}
