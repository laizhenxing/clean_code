package orm

import (
	"cc/domain/entity"
	"github.com/jinzhu/gorm"
)

type AppOrmHandler struct {
	dbConn *gorm.DB
}

func NewAppOrmHandler(db *gorm.DB) *AppOrmHandler {
	return &AppOrmHandler{dbConn: db}
}

func (a *AppOrmHandler) FindAll() ([]*entity.App, error) {
	var apps []*entity.App
	err := a.dbConn.Find(&apps).Error
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (a *AppOrmHandler) Save(app entity.App) error {
	return a.dbConn.Create(&app).Error
}

func (a *AppOrmHandler) FindOne(id uint64) (*entity.App, error) {
	var app entity.App
	err := a.dbConn.Where("id=?", id).Find(&app).Error
	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (a *AppOrmHandler) Delete(id uint64) error {
	return a.dbConn.Where("id=?", id).Delete(&entity.App{}).Error
}

func (a *AppOrmHandler) Update(app entity.App) error {
	return a.dbConn.Model(&entity.App{}).Where("id=?", app.ID).Updates(app).Error
}

