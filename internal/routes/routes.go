package routes

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	r.Mount("/api/v1/todos", TodoRoutes())
	r.Mount("/api/v1/users", UserRoutes())
	// 새로운 라우트를 추가하려면 여기서 r.Mount()를 호출합니다.
}
