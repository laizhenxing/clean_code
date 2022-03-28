package application

import (
	"cc/domain/entity"
	"cc/domain/repository"
)

type  UserInteractor struct {
	userRepo repository.UserRepository
}

func NewUserInteractor(up repository.UserRepository) *UserInteractor {
	return &UserInteractor{up}
}

func (u *UserInteractor) Get(id uint64) (*entity.User, error) {
	return u.userRepo.Get(id)
}

func (u *UserInteractor) GetAll() ([]*entity.User, error) {
	return u.userRepo.GetAll()
}

func (u *UserInteractor) CreateUser(user entity.User) error {
	return u.userRepo.Create(user)
}

func (u *UserInteractor) Update(user entity.User) error {
	return u.userRepo.Modify(user)
}

func (u *UserInteractor) Delete(id uint64) error {
	return u.userRepo.Delete(id)
}