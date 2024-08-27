package routes

import (
	"github.com/go-chi/chi/v5"
	gochi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	"time"
)

func Router() *chi.Mux {

	r := chi.NewRouter()

	applyCorsMiddleware(r)
	applyStandardMiddleware(r)
	setupRegisterRoutes(r)

	return r
}

// Middleware 순서 중요
func applyStandardMiddleware(r *chi.Mux) {
	r.Use(gochi_middleware.RealIP)
	r.Use(gochi_middleware.Logger)
	r.Use(gochi_middleware.Recoverer)
	r.Use(gochi_middleware.Timeout(60 * time.Second))
	r.Use(gochi_middleware.Throttle(100))
	r.Use(gochi_middleware.Compress(5))
	r.Use(gochi_middleware.AllowContentEncoding("application/json", "application/x-www-form-urlencoded"))
	r.Use(gochi_middleware.CleanPath)
	r.Use(gochi_middleware.RedirectSlashes)
}

func applyCorsMiddleware(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

func setupRegisterRoutes(r *chi.Mux) {
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Mount("/api/v1/todos", TodoRoutes())
	r.Mount("/api/v1/users", UserRoutes())
	// 새로운 라우트를 추가하려면 여기서 r.Mount()를 호출합니다.
}
