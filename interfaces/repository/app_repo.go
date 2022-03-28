package repository

import (
	"cc/domain/entity"
)

type AppHandler interface {
	FindAll() ([]*entity.App, error)
	Save(app entity.App) error
	FindOne(uint64) (*entity.App, error)
	Delete(uint64) error
	Update(app entity.App) error
}

type AppRepo struct {
	handler AppHandler
}

func NewAppRepo(handler AppHandler) *AppRepo {
	return &AppRepo{handler}
}

func (repo *AppRepo) SaveApp(app entity.App) error {
	err := repo.handler.Save(app)
	if err != nil {
		return err
	}

	return nil
}

func (repo *AppRepo) FindOne(id uint64) (*entity.App, error) {
	return repo.handler.FindOne(id)
}

func (repo *AppRepo) FindAll() ([]*entity.App, error) {
	results, err := repo.handler.FindAll()
	if err != nil {
		return results, err
	}
	return results, nil
}

func (repo *AppRepo) Delete(id uint64) error {
	return repo.handler.Delete(id)
}

func (repo *AppRepo) Update(app entity.App) error {
	return repo.handler.Update(app)
}