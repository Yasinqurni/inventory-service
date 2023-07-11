package service

import (
	"github.com/Yasinqurni/be-project/src/app/user/model"
	userRepository "github.com/Yasinqurni/be-project/src/app/user/repository"
)

type UserService interface {
	GetByID(id uint) (*model.User, error)
	GetByIDs(ids string) (*[]model.User, error)
}

type userServiceImpl struct {
	userRepository userRepository.UserRepository
}

func NewUserServiceImpl(userRepository userRepository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func (s *userServiceImpl) GetByID(id uint) (*model.User, error) {

	response, err := s.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *userServiceImpl) GetByIDs(ids string) (*[]model.User, error) {

	response, err := s.userRepository.GetByIDs(ids)
	if err != nil {
		return nil, err
	}
	return response, nil
}
