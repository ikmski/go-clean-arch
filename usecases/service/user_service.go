package service

import (
	"github.com/ikmski/go-clean-arch/entity"
	"github.com/ikmski/go-clean-arch/usecases/presenter"
	"github.com/ikmski/go-clean-arch/usecases/repository"
)

type userService struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserService interface {
	Create(u *entity.User) (int64, error)
	Get(id int64) (*entity.User, error)
}

func NewUserService(repo repository.UserRepository, pre presenter.UserPresenter) UserService {
	return &userService{
		UserRepository: repo,
		UserPresenter:  pre,
	}
}

func (s *userService) Create(u *entity.User) (int64, error) {
	return s.UserRepository.Store(u)
}

func (s *userService) Get(id int64) (*entity.User, error) {
	u, err := s.UserRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return s.UserPresenter.ResponseUsers(u), nil
}
