package repository

import (
	"cc/domain/entity"
	"cc/infrastructure/utils/identity"
)

type UserHandler interface {
	Create(user entity.User) error
	Get(id uint64) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Delete(id uint64) error
	Update(user entity.User) error
}

type UserRepo struct {
	handler UserHandler
	token   identity.Authable
}

func NewUserRepo(handler UserHandler, token identity.Authable) *UserRepo {
	return &UserRepo{
		handler: handler,
		token:   token,
	}
}

func (u *UserRepo) Create(user entity.User) error {
	return u.handler.Create(user)
}

func (u *UserRepo) Get(id uint64) (*entity.User, error) {
	return u.handler.Get(id)
}

func (u *UserRepo) GetAll() ([]*entity.User, error) {
	return u.handler.GetAll()
}

func (u *UserRepo) Delete(id uint64) error {
	return u.handler.Delete(id)
}

func (u *UserRepo) Modify(user entity.User) error {
	return u.handler.Update(user)
}
