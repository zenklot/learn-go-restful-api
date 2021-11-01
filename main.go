package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zenklot/learn-go-restful-api/app"
	"github.com/zenklot/learn-go-restful-api/controller"
	"github.com/zenklot/learn-go-restful-api/exception"
	"github.com/zenklot/learn-go-restful-api/helper"
	"github.com/zenklot/learn-go-restful-api/middleware"
	"github.com/zenklot/learn-go-restful-api/repository"
	"github.com/zenklot/learn-go-restful-api/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
