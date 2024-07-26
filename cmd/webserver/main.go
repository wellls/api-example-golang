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
	"github.com/wellls/api-example-golang/internal/handler"
	"github.com/wellls/api-example-golang/internal/handler/routes"
	"github.com/wellls/api-example-golang/internal/repository/categoryrepository"
	"github.com/wellls/api-example-golang/internal/repository/productrepository"
	"github.com/wellls/api-example-golang/internal/repository/userrepository"
	"github.com/wellls/api-example-golang/internal/service/categoryservice"
	"github.com/wellls/api-example-golang/internal/service/productservice"
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

	queries := sqlc.New(dbConnection)

	// user
	userRepo := userrepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(userRepo)

	// category
	categoryRepo := categoryrepository.NewCategoryRepository(dbConnection, queries)
	newCategoryService := categoryservice.NewCategoryService(categoryRepo)

	// product
	productRepo := productrepository.NewProductRepository(dbConnection, queries)
	productsService := productservice.NewProductService(productRepo)

	newHandler := handler.NewHandler(newUserService, newCategoryService, productsService)

	// init routes
	router := chi.NewRouter()
	routes.InitRoutes(router, newHandler)
	routes.InitDocsRoutes(router)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}
}
