package controller

import (
	"github.com/ikmski/go-clean-arch/entity"
	"github.com/ikmski/go-clean-arch/usecases/service"
)

type userController struct {
	userService service.UserService
}

type UserController interface {
	CreateUser(user *entity.User) (int64, error)
	GetUser(int64) (*entity.User, error)
}

func NewUserController(s service.UserService) UserController {
	return &userController{s}
}

func (c *userController) CreateUser(user *entity.User) (int64, error) {
	return c.userService.Create(user)
}

func (c *userController) GetUser(id int64) (*entity.User, error) {
	u := &entity.User{}
	u, err := c.userService.Get(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
