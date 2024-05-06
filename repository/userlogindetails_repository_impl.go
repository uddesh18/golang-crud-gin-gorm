package repository

import (
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type UserLoginDetailsRepositoryImpl struct {
	Db *gorm.DB
}

// FetchData implements UserLoginDetailsRepository.
func (u *UserLoginDetailsRepositoryImpl) FetchData() []model.UserLoginDetails {
	var userlogin []model.UserLoginDetails
	result := u.Db.Find(&userlogin)
	helper.ErrorPanic(result.Error)
	return userlogin
}

func NewUserLoginDetailsRepositoryImpl(Db *gorm.DB) UserLoginDetailsRepository {
	return &UserLoginDetailsRepositoryImpl{Db: Db}
}
