package services

import (
	"DAY_4/models"
	"DAY_4/transport"
)

func (c *services) CreateBook(book *models.Book) (*transport.Response, error) {
	bookMapping := &models.Book{
		Title:  book.Title,
		Writer: book.Writer,
	}

	err := c.repo.CreateBook(bookMapping)
	if err != nil {
		return nil, err
	}

	result := &transport.Response{
		Code:    201,
		Status:  "success",
		Message: "Success to create book",
		Data:    err,
	}

	return result, err
}

func (c *services) UpdateBook(book *models.Book, id int) (*transport.Response, error) {
	bookMapping := &models.Book{
		Title:  book.Title,
		Writer: book.Writer,
	}

	err := c.repo.UpdateBook(bookMapping, id)
	if err != nil {
		return nil, err
	}

	result := &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Success to update book",
		Data:    err,
	}

	return result, nil
}

func (c *services) DeleteBook(id int) (*transport.Response, error) {
	err := c.repo.DeleteBook(id)
	if err != nil {
		return nil, err
	}

	result := &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Success to delete book",
		Data:    err,
	}

	return result, nil
}

func (c *services) GetBookById(id int) (*transport.Response, error) {
	book, err := c.repo.GetBookById(id)
	if err != nil {
		return nil, err
	}

	bookMapping := &transport.BookMapping{
		Title:     book.Title,
		Writer:    book.Writer,
		ID:        book.ID,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}

	result := &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Success get book",
		Data:    bookMapping,
	}
	return result, nil
}

func (c *services) GetAllBooks(keywords string) (*transport.Response, error) {
	books, err := c.repo.GetAllBooks(keywords)
	if err != nil {
		return nil, err
	}

	var arrBooks []transport.BookMapping
	for _, book := range books {
		bookMapping := transport.BookMapping{
			Title:     book.Title,
			Writer:    book.Writer,
			ID:        book.ID,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
		}

		arrBooks = append(arrBooks, bookMapping)
	}

	result := &transport.Response{
		Code:    200,
		Status:  "success",
		Message: "Success get all books",
		Data:    arrBooks,
	}
	return result, nil
}
