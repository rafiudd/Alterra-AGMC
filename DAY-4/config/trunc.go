package config

import (
	"log"

	"gorm.io/gorm"
)

type trunc struct {
	DB *gorm.DB
}

func NewTrunc() *trunc {
	return &trunc{GetQuery()}
}

func (t *trunc) DeleteDataBooks() {
	log.Println("success delete data books")
	t.DB.Exec("DELETE FROM books")
}

func (t *trunc) DeleteDataUsers() {
	log.Println("success delete data users")
	t.DB.Exec("DELETE FROM users")
}
