package handlers

import (
	"DAY_4/models"
	"DAY_4/transport"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) CreateBook(c echo.Context) error {
	book := new(models.Book)
	c.Bind(book)
	response := new(transport.Response)
	result, err := h.service.CreateBook(book)

	if err != nil {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to create book"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data
	return c.JSON(http.StatusCreated, response)
}

func (h *handler) UpdateBook(c echo.Context) error {
	book := new(models.Book)
	c.Bind(book)
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(transport.Response)
	result, err := h.service.UpdateBook(book, idInt)

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Failed to update book"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data
	return c.JSON(http.StatusOK, response)
}

func (h *handler) DeleteBook(c echo.Context) error {
	book := new(models.Book)
	c.Bind(book)
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(transport.Response)
	result, err := h.service.DeleteBook(idInt)

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Failed to delete book"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetBookById(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(transport.Response)
	result, err := h.service.GetBookById(idInt)

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Book not found"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data

	return c.JSON(http.StatusOK, response)
}

func (h *handler) GetAllBooks(c echo.Context) error {
	response := new(transport.Response)
	result, err := h.service.GetAllBooks(c.QueryParam("keywords"))

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Books not found"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Data = result.Data
	return c.JSON(http.StatusOK, response)
}
