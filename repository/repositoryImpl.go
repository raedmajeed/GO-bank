package repository

import (
	"fmt"
	"raedmajeed/model"
	"raedmajeed/repository/interfaces"

	"gorm.io/gorm"
)

type RepositoryImpl struct {
	DB *gorm.DB
}

// Create implements repository.
func (s *RepositoryImpl) Create(user model.Entity) error {
	fmt.Println("reached repositroy =====.>>.........")
	s.DB.Create(&user)

	return nil
}

// FindAll implements repository.
func (s *RepositoryImpl) FindAll(user model.Entity) error {
	fmt.Println("FindAll Implemented")
	return nil
}

func NewRepositoryImpl(DB *gorm.DB) interfaces.Repository {
	return &RepositoryImpl{
		DB: DB,
	}
}
