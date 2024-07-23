package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wellls/api-example-golang/config/env"
	"github.com/wellls/api-example-golang/config/logger"
	"github.com/wellls/api-example-golang/internal/database"
	"github.com/wellls/api-example-golang/internal/database/sqlc"
	"github.com/wellls/api-example-golang/internal/handler/routes"
	"github.com/wellls/api-example-golang/internal/handler/userhandler"
	"github.com/wellls/api-example-golang/internal/repository/userrepository"
	"github.com/wellls/api-example-golang/internal/service/userservice"
)

func main() {
	logger.InitLogger()
	slog.Info("starting api")

	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failed to load environment variables", err, slog.String("package", "main"))
		return
	}
	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
		return
	}

	router := chi.NewRouter()
	queries := sqlc.New(dbConnection)

	// user
	userRepo := userrepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(userRepo)
	newUserHandler := userhandler.NewUserHandler(newUserService)

	// init routes
	routes.InitUserRoutes(router, newUserHandler)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}
}
