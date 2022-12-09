package repository

import (
	"context"
	"gorm.io/gorm"
	"pergipergi/model"
)

type UserRepository interface {
	AddUser(ctx context.Context, user model.User) (model.User, error)
	GetUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	DeleteUser(ctx context.Context, ID int) error
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) AddUser(ctx context.Context, user model.User) (model.User, error) {
	err := u.db.WithContext(ctx).Model(&model.User{}).Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := u.db.WithContext(ctx).Model(&model.User{}).Select("*").Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	err := u.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRepository) DeleteUser(ctx context.Context, ID int) error {
	err := u.db.WithContext(ctx).Where("id = ?", ID).Delete(model.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).Select("*").Where("username = ?", username).Scan(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
