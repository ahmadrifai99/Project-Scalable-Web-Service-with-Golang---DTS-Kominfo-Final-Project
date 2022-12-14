package service

import (
	"final-project/configuration"
	"final-project/dto"
	"final-project/helper"
	"final-project/interfaces"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Conf configuration.Config
	repo interfaces.UserRepository
}

func NewUserService(conf configuration.Config, repo interfaces.UserRepository) interfaces.UserService {
	return &UserService{
		Conf: conf,
		repo: repo,
	}
}

func (s *UserService) Register(user *dto.Register) (dto.Register, error) {
	// bcrypt password
	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	// set password
	user.Password = string(pass)

	return s.repo.Register(user)
}

func (s *UserService) Login(email, password string) (string, error) {
	login, err := s.repo.Login(email, password)
	if err != nil {
		return "", err
	}
	_ = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(password))

	token, _ := helper.GenerateAccessToken(login.ID, login.Email)

	return token, nil
}

func (s *UserService) UpdateUser(id uint, user *dto.UpdateUser) (dto.UpdateUser, error) {
	return s.repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
