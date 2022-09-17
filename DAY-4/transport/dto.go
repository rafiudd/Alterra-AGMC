package transport

import (
	"time"
)

type (
	Response struct {
		Code    int         `json:"code"`
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	UserMapping struct {
		ID        int       `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	BookMapping struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Writer    string    `json:"writer"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	LoginResponse struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}
)
