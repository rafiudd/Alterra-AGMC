package lib

import (
	"dynamic-crud/models"
)

func (r *repositories) CreateBook(book *models.Book) error {
	if err := r.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositories) UpdateBook(book *models.Book, id int) error {
	if err := r.db.Model(book).Where("id = ?", id).Updates(book).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositories) DeleteBook(id int) error {
	var book *models.Book
	if err := r.db.Delete(&book, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *repositories) GetBookById(id int) (models.Book, error) {
	var book models.Book
	result := r.db.Where("id = ?", id).First(&book)
	return book, result.Error
}

func (r *repositories) GetAllBooks(keywords string) ([]models.Book, error) {
	var books []models.Book
	result := r.db.Where("title LIKE ? OR writer LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&books)

	return books, result.Error
}
