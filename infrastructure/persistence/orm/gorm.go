package orm

import (
	"cc/domain/entity"
	"github.com/jinzhu/gorm"
)

type Gorm struct {
	conn *gorm.DB
}

var ORM *Gorm

func NewGorm(driver, dsn string) (*Gorm, error) {
	conn, err := gorm.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	ORM = &Gorm{conn: conn}
	return ORM, nil
}

func (g *Gorm) GetConn() *gorm.DB {
	return g.conn
}

func (g *Gorm) Close() error {
	return g.conn.Close()
}

func (g *Gorm) AutoMigrate()  {
	g.conn.AutoMigrate(&entity.App{})
	g.conn.AutoMigrate(&entity.User{})
}