package dto

import (
	"time"
	"todo-api-golang/ent"
)

// TodoDTO is a Data Transfer Object for Todo entity.
type TodoDTO struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

// ConvertTodoToDTO converts a Todo entity to TodoDTO.
func ConvertTodoToDTO(todo *ent.Todo) TodoDTO {
	return TodoDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      string(todo.Status),
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		DeletedAt:   todo.DeletedAt,
	}
}
