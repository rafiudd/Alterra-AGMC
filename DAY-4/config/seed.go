package config

import (
	"DAY_4/models"
	"DAY_4/utils"
	"log"

	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeed() *seed {
	return &seed{GetQuery()}
}

func (s *seed) UsersSeeder() {
	hash, _ := utils.HashPassword("12345")
	var users = []models.User{
		{
			Username: "test1",
			Email:    "test1@mail.com",
			Password: hash,
		},
		{
			Username: "test2",
			Email:    "test2@mail.com",
			Password: hash,
		},
	}

	if err := s.DB.Create(&users).Error; err != nil {
		log.Printf("cannot seed data users, with error %v\n", err)
	}
	log.Println("success seed data users")
}

func (s *seed) BooksSeeder() {
	var books = []models.Book{
		{
			Title:  "Test Book 1",
			Writer: "John Doe",
		},
		{
			Title:  "Test Book 2",
			Writer: "Jane Doe",
		},
	}

	if err := s.DB.Create(&books).Error; err != nil {
		log.Printf("cannot seed data books, with error %v\n", err)
	}
	log.Println("success seed data books")
}
