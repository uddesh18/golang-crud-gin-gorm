package service

import (
	"fmt"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationServiceImpl struct {
	UserRegistrationRepository repository.UserRegistrationRepository
	Validate                   *validator.Validate
}

// UserSave implements UserRegistrationService.
func (u *UserRegistrationServiceImpl) UserSave(usersave request.UserRegistrationRequest) {
	err := u.Validate.Struct(usersave)
	helper.ErrorPanic(err)
	fmt.Println(usersave.Username)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(usersave.Password), bcrypt.DefaultCost)
	userRegModel := model.UserRegistration{
		Username: usersave.Username,
		Password: string(hashedPassword),
		Email:    usersave.Email,
		Role:     usersave.Role,
	}
	u.UserRegistrationRepository.Save(userRegModel)
}

func NewUserRegistrationServiceImpl(userregistrationrepository repository.UserRegistrationRepository, validate *validator.Validate) UserRegistrationService {
	return &UserRegistrationServiceImpl{
		UserRegistrationRepository: userregistrationrepository,
		Validate:                   validate,
	}
}
