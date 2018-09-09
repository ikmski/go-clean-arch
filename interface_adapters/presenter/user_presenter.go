package presenter

import "github.com/ikmski/go-clean-arch/entity"

type userPresenter struct {
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

type UserPresenter interface {
	ResponseUsers(u *entity.User) *entity.User
}

func (p *userPresenter) ResponseUsers(u *entity.User) *entity.User {
	return u
}
