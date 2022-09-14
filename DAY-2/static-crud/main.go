package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	v1 := e.Group("/v1")
	books := v1.Group("/books")

	books.GET("", getAllBooks)
	books.GET("/:id", getBookById)
	books.POST("", createNewBook)
	books.PUT("/:id", updateBookById)
	books.DELETE("/:id", deleteBookById)

	e.Logger.Fatal(e.Start(":8080"))
}

// Book format data
type Book struct {
	ID       uint   `json:"id,omitempty"`
	BookName string `json:"book_name,omitempty"`
	Author   string `json:"author,omitempty"`
	Pages    int    `json:"pages,omitempty"`
	Slug     string `json:"slug,omitempty"`
}

// HTTPResponse format
type HTTPResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var books = []Book{
	{ID: 1, BookName: "Cantik itu Luka", Author: "Eka Kurniawan", Pages: 150, Slug: "cantik-itu-luka"},
	{ID: 2, BookName: "Bumi Manusia", Author: "Pramoedya Ananta Toer", Pages: 50, Slug: "bumi-manusia "},
	{ID: 2, BookName: "Negeri 5 Menara", Author: "Ahmad Fuadi", Pages: 50, Slug: "negeri-5-menara"},
}

func wrapperResponse(code int, message string, response interface{}) *HTTPResponse {
	commonResponse := new(HTTPResponse)
	commonResponse.Code = code
	commonResponse.Message = message
	commonResponse.Data = response
	return commonResponse
}

func (resp *HTTPResponse) JSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	return json.NewEncoder(w).Encode(resp)
}

func getBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, book := range books {
		if book.ID == uint(id) {
			return wrapperResponse(http.StatusOK, "success get detail", book).JSON(c.Response())
		}
	}

	return wrapperResponse(http.StatusBadRequest, "failed get detail", &Book{}).JSON(c.Response())
}

func getAllBooks(c echo.Context) error {
	return wrapperResponse(http.StatusOK, "success get all books", books).JSON(c.Response())
}

func createNewBook(c echo.Context) error {
	book := Book{}
	err := json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "failed create book", &Book{}).JSON(c.Response())
	}

	books = append(books, book)
	fmt.Println("Current Repo:", books)

	return wrapperResponse(http.StatusCreated, "success create book", &Book{}).JSON(c.Response())
}

func updateBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookPayload := Book{}
	err := json.NewDecoder(c.Request().Body).Decode(&bookPayload)
	if err != nil {
		return wrapperResponse(http.StatusBadRequest, "failed update book", &Book{}).JSON(c.Response())
	}
	for i, book := range books {
		if book.ID == uint(id) {
			books[i] = bookPayload
			return wrapperResponse(http.StatusOK, "success update book", &Book{}).JSON(c.Response())
		}
	}

	return wrapperResponse(http.StatusBadRequest, "failed update book", &Book{}).JSON(c.Response())
}

func deleteBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, book := range books {
		if book.ID == uint(id) {
			books[i] = books[0]
			books = books[1:]
			fmt.Println("Current Repo:", books)
			return wrapperResponse(http.StatusOK, "success delete book", &Book{}).JSON(c.Response())
		}
	}

	return wrapperResponse(http.StatusBadRequest, "failed delete book", &Book{}).JSON(c.Response())
}
