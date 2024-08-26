package routes

import (
	"github.com/go-chi/chi/v5"
	"todo-api-golang/internal/handlers"
)

func TodoRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handlers.CreateTodo)
	r.Get("/", handlers.ListTodos)
	r.Get("/{id}", handlers.GetTodo)
	r.Put("/{id}", handlers.UpdateTodo)
	r.Delete("/{id}", handlers.DeleteTodo)

	return r
}
