package routes

import (
	"github.com/go-chi/chi/v5"
	"todo-api-golang/edge/database"
	"todo-api-golang/internal/handlers"
	"todo-api-golang/internal/service"
)

func TodoRoutes() chi.Router {
	r := chi.NewRouter()

	client := database.InitDB()

	todoService := service.NewTodoService(client)
	todoHandlers := handlers.NewTodoHandler(todoService)

	r.Post("/", todoHandlers.CreateTodo)
	r.Get("/", todoHandlers.ListTodos)
	r.Get("/{id}", todoHandlers.GetTodo)
	r.Put("/{id}", todoHandlers.UpdateTodo)
	r.Delete("/{id}", todoHandlers.DeleteTodo)
	r.Put("/{id}/status", todoHandlers.UpdateTodoStatus)

	return r
}
