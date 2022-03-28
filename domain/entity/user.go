package entity

import "time"

type User struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment" bson:"id"`
	Username  string    `json:"username" gorm:"username" bson:"username"`
	Password  string    `json:"password" gorm:"password" bson:"password"`
	Phone     string    `json:"phone" gorm:"phone" bson:"phone"`
	Email     string    `json:"email" gorm:"email" bson:"phone"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at;default:CURRENT_TIMESTAMP" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at;default:CURRENT_TIMESTAMP" bson:"updated_at"`
}
