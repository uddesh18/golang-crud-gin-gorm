package model

import "time"

type UserLoginDetails struct {
	Id          int       `gorm:"type:int;primary_key"`
	Username    string    `gorm:"size:255;not null;unique" json:"username"`
	Password    string    `gorm:"size:255;not null;" json:"password"`
	AccessToken string    `gorm:"size:5000;not null;" json:"access_token"`
	CreatedDate time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_date"`
	IsActive    bool      `gorm:"type:bool" json:"is_active"`
}
