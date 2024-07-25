package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/wellls/api-example-golang/config/env"
	"github.com/wellls/api-example-golang/internal/handler/middleware"
	"github.com/wellls/api-example-golang/internal/handler/userhandler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Use(middleware.LoggerData)

	router.Post("/user", h.CreateUser)
	router.Route("/user", func(r chi.Router) {
		r.Use(jwtauth.Verifier(env.Env.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Patch("/", h.UpdateUser)
		r.Get("/", h.GetUserByID)
		r.Delete("/", h.DeleteUser)
		r.Get("/list-all", h.FindManyUsers)
		r.Patch("/password", h.UpdateUserPassword)
	})
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
	})
}
