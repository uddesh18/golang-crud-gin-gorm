package repository

import (
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type UserLoginDetailsRepositoryImpl struct {
	Db *gorm.DB
}

// FetchData implements UserLoginDetailsRepository.
func (u *UserLoginDetailsRepositoryImpl) FetchData(user model.UserLoginDetails) {
	panic("unimplemented")
}

func NewUserLoginDetailsRepositoryImpl(Db *gorm.DB) UserLoginDetailsRepository {
	return &UserLoginDetailsRepositoryImpl{Db: Db}
}
