package handlers

import (
	"DAY_3/controllers"

	"github.com/labstack/echo/v4"
)

type handler struct {
	controller controllers.Controllers
}

type Handlers interface {
	//! Health Check Handler
	HealthCheck(c echo.Context) error

	//! User Handler
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	GetUserById(c echo.Context) error
	GetAllUsers(c echo.Context) error
	DeleteUser(c echo.Context) error
	UserLogin(c echo.Context) error

	//! Book Handler
	CreateBook(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
	GetBookById(c echo.Context) error
	GetAllBooks(c echo.Context) error
}

func NewHandlers(controller controllers.Controllers) Handlers {
	return &handler{controller: controller}
}
