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
func (u *UserLoginDetailServiceImpl) LoginUser() []response.UldResponse {
	result := u.UserLoginDetailsRepository.FetchData()

	var tags []response.UldResponse
	for _, value := range result {
		tag := response.UldResponse{
			Id:           value.Id,
			RegisteredId: value.RegisteredId,
			Username:     value.Username,
			AccessToken:  value.AccessToken,
			Password:     value.Password,
			CreatedDate:  value.CreatedDate.String(),
			IsActive:     value.IsActive,
		}
		tags = append(tags, tag)
	}

	return tags
}

func NewUserLoginDetailServiceImpl(userlogindetailsRepository repository.UserLoginDetailsRepository, validate *validator.Validate) UserLoginDetailService {
	return &UserLoginDetailServiceImpl{
		UserLoginDetailsRepository: userlogindetailsRepository,
		Validate:                   validate,
	}
}
