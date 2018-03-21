package model

// Todo - Todo Item
type Todo struct {
	ID          uint64 `json:"todoId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
