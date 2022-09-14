package routes

import (
	"dynamic-crud/config"
	"dynamic-crud/controllers"
	"dynamic-crud/handlers"
	lib "dynamic-crud/lib/repositories"

	"github.com/labstack/echo/v4"
)

func Routes() {
	e := echo.New()

	repository := lib.NewRepositories(config.GetQuery())
	controller := controllers.NewControllers(repository)
	handler := handlers.NewHandlers(controller)

	e.GET("/v1/healthcheck", handler.HealthCheck)

	//! User routes
	e.POST("/v1/users", handler.CreateUser)
	e.PUT("/v1/users/:id", handler.UpdateUser)
	e.GET("/v1/users/:id", handler.GetUserById)
	e.GET("/v1/users", handler.GetAllUsers)
	e.DELETE("/v1/users/:id", handler.DeleteUser)

	//! Book routes
	e.POST("/v1/books", handler.CreateBook)
	e.GET("/v1/books/:id", handler.GetBookById)
	e.GET("/v1/books", handler.GetAllBooks)
	e.PUT("/v1/books/:id", handler.UpdateBook)
	e.DELETE("/v1/books/:id", handler.DeleteBook)

	e.Logger.Fatal(e.Start(":8080"))
}
