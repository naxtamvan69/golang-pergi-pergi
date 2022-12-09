package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"pergipergi/model"
	"pergipergi/repository"
)

type UserService interface {
	AddUserService(ctx context.Context, user model.User) (model.User, error)
	GetUsersService(ctx context.Context) ([]model.User, error)
	UpdateUserService(ctx context.Context, user model.User) (model.User, error)
	DeleteUserService(ctx context.Context, ID int) error
	GetUserByUsernameService(ctx context.Context, username string) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (u *userService) AddUserService(ctx context.Context, user model.User) (model.User, error) {
	_, err := u.userRepository.GetUserByUsername(ctx, user.Username)
	if err == nil {
		return model.User{}, errors.New("username already used")
	}

	pass, err := u.HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}
	user.Password = pass

	return u.userRepository.AddUser(ctx, user)
}

func (u *userService) GetUsersService(ctx context.Context) ([]model.User, error) {
	return u.userRepository.GetUsers(ctx)
}

func (u *userService) UpdateUserService(ctx context.Context, user model.User) (model.User, error) {
	return u.userRepository.UpdateUser(ctx, user)
}

func (u *userService) DeleteUserService(ctx context.Context, ID int) error {
	return u.userRepository.DeleteUser(ctx, ID)
}

func (u *userService) GetUserByUsernameService(ctx context.Context, username string) (model.User, error) {
	return u.userRepository.GetUserByUsername(ctx, username)
}

func (u *userService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func (u *userService) CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
