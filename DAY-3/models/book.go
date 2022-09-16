package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        int            `json:"id" form:"id"`
	Title     string         `json:"title" form:"title"`
	Writer    string         `json:"writer" form:"writer"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted" form:"deleted"`
}
