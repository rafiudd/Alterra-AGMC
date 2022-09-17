package handlers

import (
	"DAY_4/services"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service services.Services
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

func NewHandlers(service services.Services) Handlers {
	return &handler{service: service}
}
