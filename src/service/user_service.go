package service

import (
	"CNAD_CloudShop/src/domain"
	"CNAD_CloudShop/src/repository"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) UserExists(username string) bool {
	_, err := u.userRepo.Get(username)
	return err == nil
}

func (u *UserService) Register(username string) bool {
	if u.UserExists(username) {
		return false
	}
	err := u.userRepo.Create(&domain.User{Username: username})
	return err == nil
}