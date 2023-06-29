package repo

import (
	"context"

	"github.com/YianAndCode/go-web-app-skeleton/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Get(ctx context.Context, userId int) (*model.User, error) {
	user := model.User{UserId: userId}
	err := r.db.First(&user).Error
	return &user, err
}
