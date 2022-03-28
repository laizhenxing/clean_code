package repository

import "cc/domain/entity"

type AppRepository interface {
	SaveApp(app entity.App) error
	FindOne(uint64) (*entity.App, error)
	FindAll() ([]*entity.App, error)
	Update(app entity.App) error
	Delete(uint64) error
}