package application

import (
	"cc/domain/entity"
	"cc/domain/repository"
)

type AppInteractor struct {
	AppRepository repository.AppRepository
}

// 注入 AppRepository 接口
func NewAppInteractor(repository repository.AppRepository) AppInteractor {
	return AppInteractor{repository}
}

func (interactor *AppInteractor) CreateApp(app entity.App) error {
	return interactor.AppRepository.SaveApp(app)
}

func (interactor *AppInteractor) FindAll() ([]*entity.App, error) {
	results, err := interactor.AppRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (interactor *AppInteractor) FindOne(id uint64) (*entity.App, error) {
	return interactor.AppRepository.FindOne(id)
}

func (interactor *AppInteractor) Delete(id uint64) error {
	return interactor.AppRepository.Delete(id)
}

func (interactor *AppInteractor) Update(app entity.App) error {
	return interactor.AppRepository.Update(app)
}
