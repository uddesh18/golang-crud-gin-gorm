package model

import "time"

type UserLoginDetails struct {
	Id           int       `gorm:"type:int;primary_key"`
	Username     string    `gorm:"size:255;not null;unique"`
	Password     string    `gorm:"size:255;not null;" `
	AccessToken  string    `gorm:"size:5000;not null;" `
	CreatedDate  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" `
	IsActive     bool      `gorm:"type:bool" `
	RegisteredId int       `gorm:"type:int" `
}
