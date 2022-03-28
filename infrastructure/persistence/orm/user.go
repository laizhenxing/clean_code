package orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"cc/domain/entity"
	"cc/infrastructure/utils/identity"
)

type UserHandler struct {
	dbConn *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

func (u *UserHandler) Create(user entity.User) error {
	err := u.dbConn.Save(&user).Error
	if err != nil {
		return err
	}
	token, err := identity.GenToken(identity.NewClaims(user.ID, true, nil), []byte(viper.GetString("app.jwt_secret")))

	if err != nil {
		return err
	}
	fmt.Println("token: ", token)
	return nil
}

func (u *UserHandler) Get(id uint64) (*entity.User, error) {
	var user *entity.User
	err := u.dbConn.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserHandler) GetAll() ([]*entity.User, error) {
	var users []*entity.User
	err := u.dbConn.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserHandler) Delete(id uint64) error {
	return u.dbConn.Where("id=?", id).Delete(&entity.User{}).Error
}

func (u *UserHandler) Update(user entity.User) error {
	return u.dbConn.Model(&entity.User{}).Where("id=?", user.ID).Updates(user).Error
}
