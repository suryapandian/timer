package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//GetRouter - router configuration
func GetRouter() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	setTimerRoutes(mux)

	return mux
}
