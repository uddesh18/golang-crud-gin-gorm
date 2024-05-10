package repository

import (
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type UserRegistrationImpl struct {
	Db *gorm.DB
}

// Save implements UserRegistrationRepository.
func (u *UserRegistrationImpl) Save(userregistration model.UserRegistration) {
	result := u.Db.Create(&userregistration)
	helper.ErrorPanic(result.Error)
}

func NewUserRegistrationImpl(Db *gorm.DB) UserRegistrationRepository {
	return &UserRegistrationImpl{Db: Db}
}
