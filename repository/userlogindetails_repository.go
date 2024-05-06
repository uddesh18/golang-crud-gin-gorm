package repository

import "golang-crud-gin/model"

type UserLoginDetailsRepository interface {
	FetchData() []model.UserLoginDetails
}
