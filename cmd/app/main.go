package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"todo-api-golang/edge/log"
	"todo-api-golang/internal/routes"
)

func main() {
	log.Logger.Println("Starting server...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 통합된 라우트 등록
	routes.RegisterRoutes(r)

	log.Logger.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Logger.Fatal()
	}
}
