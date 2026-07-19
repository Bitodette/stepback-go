package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"stepback-golang/internal/handler"
	"stepback-golang/internal/middleware"
	"stepback-golang/internal/service"
)

func New(authHandler *handler.AuthHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.CORS)

	r.Route("/api", func(r chi.Router) {
		r.Post("/users/register", authHandler.Register)
		r.Post("/users/login", authHandler.Login)
		r.Post("/users/token/refresh", authHandler.RefreshToken)

		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
	})

	return r
}

func NewWithDeps(authService *service.AuthService) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.CORS)

	authHandler := handler.NewAuthHandler(authService)

	r.Route("/api", func(r chi.Router) {
		r.Post("/users/register", authHandler.Register)
		r.Post("/users/login", authHandler.Login)
		r.Post("/users/token/refresh", authHandler.RefreshToken)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth)

			r.Get("/users/profile", func(w http.ResponseWriter, r *http.Request) {
				// TODO: implement profile handler
				w.Write([]byte("profile"))
			})
		})

		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
	})

	return r
}
