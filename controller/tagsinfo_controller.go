package controller

import (
	"fmt"
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TagsInfoController struct {
	tagsinfo_service service.TagsInfoService
}

func NewTagsInfoController(service service.TagsInfoService) *TagsInfoController {
	return &TagsInfoController{
		tagsinfo_service: service,
	}
}

// CreateTags		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body request.CreateTagsRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (controller *TagsInfoController) CreateInfo(ctx *gin.Context) {
	log.Info().Msg("create info tags")
	createInfoTagsRequest := request.CreateTagsInfoRequest{}
	err := ctx.ShouldBindJSON(&createInfoTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsinfo_service.CreateInfo(createInfoTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tags
// @Success			200 {obejct} response.Response{}
// @Router			/tags [get]
func (controller *TagsInfoController) FindAllInfo(ctx *gin.Context) {
	log.Info().Msg("findAllInfo tags")
	tagResponse := controller.tagsinfo_service.FindAllInfo()
	fmt.Println(tagResponse)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
