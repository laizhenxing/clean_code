package repository

import "cc/domain/entity"

type UserRepository interface {
	Create(user entity.User) error
	Get(id uint64) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Delete(id uint64) error
	Modify(user entity.User) error
}
