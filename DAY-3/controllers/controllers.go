package controllers

import (
	"DAY_3/models"
	repositories "DAY_3/repositories"
	"DAY_3/transport"
)

type controllers struct {
	repo repositories.Repositories
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
	UserLogin(username, password string) (*transport.Response, error)

	// ! Health Check Controllers
	HealthCheck() *transport.Response
}

func NewControllers(r repositories.Repositories) Controllers {
	return &controllers{repo: r}
}
