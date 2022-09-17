package repositories

import (
	"DAY_4/models"

	"gorm.io/gorm"
)

type repositories struct {
	db *gorm.DB
}

type Repositories interface {
	//! Book Repositories
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book, id int) error
	DeleteBook(id int) error
	GetBookById(id int) (models.Book, error)
	GetAllBooks(keywords string) ([]models.Book, error)

	//! User Repositories
	CreateUser(user *models.User) error
	UpdateUser(user *models.User, id int) error
	DeleteUser(id int) error
	GetUserById(id int) (models.User, error)
	GetAllUsers(keywords string) ([]models.User, error)
	UserLogin(username string) (models.User, error)
}

func NewRepositories(
	q *gorm.DB,
) Repositories {
	return &repositories{
		db: q,
	}
}
