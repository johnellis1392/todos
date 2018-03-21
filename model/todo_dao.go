package model

import (
	"fmt"
)

var (
	currentID uint64
	todos     []Todo
)

func init() {
	currentID = 0
	RepoCreateTodo(Todo{Title: "Write Presentation"})
	RepoCreateTodo(Todo{Title: "Host Meetup"})
}

// RepoFindTodo - Find Todo
func RepoFindTodo(id uint64) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}

	return Todo{}
}

// RepoCreateTodo - Create Todo
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

// RepoDestroyTodo - Destroy Todo
func RepoDestroyTodo(id uint64) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of '%d' to delete", id)
}
