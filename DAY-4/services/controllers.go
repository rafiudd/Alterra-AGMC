package services

import (
	"DAY_4/models"
	repositories "DAY_4/repositories"
	"DAY_4/transport"
)

type services struct {
	repo repositories.Repositories
}

type Services interface {
	// ! Book Services
	CreateBook(book *models.Book) (*transport.Response, error)
	UpdateBook(book *models.Book, id int) (*transport.Response, error)
	DeleteBook(id int) (*transport.Response, error)
	GetBookById(id int) (*transport.Response, error)
	GetAllBooks(keywords string) (*transport.Response, error)

	// ! User Services
	CreateUser(user *models.User) (*transport.Response, error)
	UpdateUser(user *models.User, id int) (*transport.Response, error)
	DeleteUser(id int) (*transport.Response, error)
	GetUserById(id int) (*transport.Response, error)
	GetAllUsers(keywords string) (*transport.Response, error)
	UserLogin(username, password string) (*transport.Response, error)

	// ! Health Check Services
	HealthCheck() *transport.Response
}

func NewServices(r repositories.Repositories) Services {
	return &services{repo: r}
}
