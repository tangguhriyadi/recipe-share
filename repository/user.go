package repository

import (
	"context"

	"github.com/tangguhriyadi/recipe-share/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]models.Users, error)
}

type UserImplementation struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserImplementation{
		db: db,
	}
}

func (ur *UserImplementation) GetAll(ctx context.Context) ([]models.Users, error) {
	var users []models.Users
	ur.db.Find(&users)

	return users, nil
}
