package service

import "golang-crud-gin/data/request"

type UserRegRequest struct {
	Username     string `validate:"required,min=1,max=200" json:"username"`
	Password     string `validate:"required,min=1,max=200" json:"password"`
	IsActive     bool   `validate:"required,min=1,max=200" json:"is_active"`
	RegisteredId int    `validate:"required,min=1,max=200" json:"registered_id"`
}

type UserRegistrationService interface {
	UserSave(usersave request.UserRegistrationRequest)
}
