package repository

import "golang-crud-gin/model"

type UserLoginDetailsRepository interface {
	FetchData(user model.UserLoginDetails)
}
