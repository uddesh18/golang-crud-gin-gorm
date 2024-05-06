package model

type UserRegistration struct {
	Id       int    `gorm:"type:int;primary_key"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255)"`
}
