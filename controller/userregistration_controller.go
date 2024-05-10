package controller

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserRegistrationController struct {
	userregistration_service service.UserRegistrationService
}

func NewUserRegistrationController(service service.UserRegistrationService) *UserRegistrationController {
	return &UserRegistrationController{
		userregistration_service: service,
	}
}

// CreateTags		godoc
// @Summary			User Create
// @Description		Save Users data in db
// @Param			tags body request.UserRegistrationRequest true "user create"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (controller *UserRegistrationController) UserCreate(ctx *gin.Context) {
	log.Info().Msg("create tags")
	createUserReg := request.UserRegistrationRequest{}
	err := ctx.ShouldBindJSON(&createUserReg)
	helper.ErrorPanic(err)

	controller.userregistration_service.UserSave(createUserReg)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "User Signed Successfully!",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
