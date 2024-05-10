package controller

import (
	"fmt"
	"golang-crud-gin/data/response"
	"golang-crud-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserLoginDetailsController struct {
	userlogindetails_service service.UserLoginDetailService
}

func NewLoginDetailsController(service service.UserLoginDetailService) *UserLoginDetailsController {
	return &UserLoginDetailsController{
		userlogindetails_service: service,
	}
}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			user login details
// @Success			200 {obejct} response.Response{}
// @Router			/allusers [get]
func (controller *UserLoginDetailsController) UserDetails(ctx *gin.Context) {
	log.Info().Msg("All User Login Details")
	uldetailsResponse := controller.userlogindetails_service.LoginUser()
	fmt.Println(uldetailsResponse)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   uldetailsResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
