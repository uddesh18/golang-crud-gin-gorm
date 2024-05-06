package service

import (
	"golang-crud-gin/data/response"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type UserLoginDetailServiceImpl struct {
	UserLoginDetailsRepository repository.UserLoginDetailsRepository
	Validate                   *validator.Validate
}

// LoginUser implements UserLoginDetailService.
func (u *UserLoginDetailServiceImpl) LoginUser() []response.Response {
	panic("unimplemented")
}

func NewUserLoginDetailServiceImpl(userlogindetailsRepository repository.UserLoginDetailsRepository, validate *validator.Validate) UserLoginDetailService {
	return &UserLoginDetailServiceImpl{
		UserLoginDetailsRepository: userlogindetailsRepository,
		Validate:                   validate,
	}
}
