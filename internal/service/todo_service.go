package service

import (
	"context"
	"time"
	"todo-api-golang/ent"
	"todo-api-golang/ent/todo"
	"todo-api-golang/internal/dto"
)

// TodoService defines the interface for Todo operations and implements it.
type TodoService interface {
	CreateTodo(ctx context.Context, form dto.TodoForm) (*dto.TodoDTO, error)
	GetTodoByID(ctx context.Context, id int) (*dto.TodoDTO, error)
	UpdateTodo(ctx context.Context, id int, form dto.TodoForm) (*dto.TodoDTO, error)
	UpdateTodoStatus(ctx context.Context, id int, status string) (*dto.TodoDTO, error)
	DeleteTodo(ctx context.Context, id int) error
	ListTodos(ctx context.Context) ([]dto.TodoDTO, error)
}

// todoService is the concrete implementation of TodoService.
type todoService struct {
	client *ent.Client
}

// NewTodoService creates a new instance of todoService.
func NewTodoService(client *ent.Client) TodoService {
	return &todoService{client: client}
}

// Implement the methods defined in the TodoService interface.

func (s *todoService) CreateTodo(ctx context.Context, form dto.TodoForm) (*dto.TodoDTO, error) {
	todoItem, err := s.client.Todo.Create().
		SetTitle(form.Title).
		SetDescription(form.Description).
		SetStatus(todo.Status(form.Status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	todoDTO := dto.ConvertTodoToDTO(todoItem)
	return &todoDTO, nil
}

func (s *todoService) GetTodoByID(ctx context.Context, id int) (*dto.TodoDTO, error) {
	todoItem, err := s.client.Todo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	todoDTO := dto.ConvertTodoToDTO(todoItem)
	return &todoDTO, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, id int, form dto.TodoForm) (*dto.TodoDTO, error) {
	todoItem, err := s.client.Todo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	_, err = todoItem.Update().
		SetTitle(form.Title).
		SetDescription(form.Description).
		SetStatus(todo.Status(form.Status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	todoDTO := dto.ConvertTodoToDTO(todoItem)
	return &todoDTO, nil
}

func (s *todoService) UpdateTodoStatus(ctx context.Context, id int, status string) (*dto.TodoDTO, error) {
	todoItem, err := s.client.Todo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	_, err = todoItem.Update().
		SetStatus(todo.Status(status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	todoDTO := dto.ConvertTodoToDTO(todoItem)
	return &todoDTO, nil
}

func (s *todoService) DeleteTodo(ctx context.Context, id int) error {
	_, err := s.client.Todo.UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return err
}

func (s *todoService) ListTodos(ctx context.Context) ([]dto.TodoDTO, error) {
	todos, err := s.client.Todo.Query().
		Where(todo.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	todoDTOs := make([]dto.TodoDTO, len(todos))
	for i, t := range todos {
		todoDTOs[i] = dto.ConvertTodoToDTO(t)
	}
	return todoDTOs, nil
}
