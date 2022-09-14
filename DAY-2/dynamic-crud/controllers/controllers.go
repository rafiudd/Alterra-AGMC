package controllers

import (
	lib "dynamic-crud/lib/repositories"
	"dynamic-crud/models"
	"dynamic-crud/transport"
)

type controllers struct {
	repo lib.Repositories
}

type Controllers interface {
	// ! Book Controllers
	CreateBook(book *models.Book) (*transport.Response, error)
	UpdateBook(book *models.Book, id int) (*transport.Response, error)
	DeleteBook(id int) (*transport.Response, error)
	GetBookById(id int) (*transport.Response, error)
	GetAllBooks(keywords string) (*transport.Response, error)

	// ! User Controllers
	CreateUser(user *models.User) (*transport.Response, error)
	UpdateUser(user *models.User, id int) (*transport.Response, error)
	DeleteUser(id int) (*transport.Response, error)
	GetUserById(id int) (*transport.Response, error)
	GetAllUsers(keywords string) (*transport.Response, error)

	// ! Health Check Controllers
	HealthCheck() *transport.Response
}

func NewControllers(r lib.Repositories) Controllers {
	return &controllers{repo: r}
}
