package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int            `json:"id" form:"id"`
	Name      string         `json:"name" form:"name"`
	Email     string         `json:"email" form:"email"`
	Password  string         `json:"password" form:"password"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted" form:"deleted"`
}
