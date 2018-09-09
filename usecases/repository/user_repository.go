package repository

import "github.com/ikmski/go-clean-arch/entity"

type UserRepository interface {
	Store(*entity.User) error
	Find(int64) (*entity.User, error)
	Remove(int64) (*entity.User, error)
}
