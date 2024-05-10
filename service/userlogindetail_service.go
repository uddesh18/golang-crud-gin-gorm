package service

import "golang-crud-gin/data/response"

type UserLoginDetailService interface {
	LoginUser() []response.UldResponse
}
