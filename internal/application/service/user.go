package service

import (
	"gorm-terminal/internal/domain/repository"

	"gorm-terminal/internal/domain/model"
)

type UserService struct {
	userRepository repository.UserRepository

}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository}
}


func (s *UserService) FindAll() ([]model.User, error) {

	users, err := s.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}


func (s *UserService) FindByID(id uint) (*model.User, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}


func (s *UserService) Create(user *model.User) error {
	if err := s.userRepository.Create(user); err != nil {
		return err
	}
	return nil
}


func (s *UserService) Update(user *model.User,id uint) error {
	if err := s.userRepository.Update(user,id); err != nil {
		return err
	}
	return nil
}


func (s *UserService) Delete(id uint) error {
	if err := s.userRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

