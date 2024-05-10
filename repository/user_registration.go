package repository

import "golang-crud-gin/model"

type UserRegistrationRepository interface {
	Save(userregistration model.UserRegistration)
}
