package presenter

import "github.com/ikmski/go-clean-arch/entity"

type UserPresenter interface {
	ResponseUsers(u *entity.User) *entity.User
}
