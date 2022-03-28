package entity

import "time"

//type App struct {
//	ID        uint64    `json:"id"`
//	Name      string    `json:"name"`
//	Version   string    `json:"version"`
//	CreatedAt time.Time `json:"created_at"`
//	UpdatedAt time.Time `json:"updated_at"`
//}

// 使用 gorm
type App struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id" bson:"id"`
	Name      string    `gorm:"name" json:"name" bson:"name"`
	Version   string    `gorm:"version" json:"version" bson:"version"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at" bson:"updated_at"`
}