package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type TagsInfoService interface {
	CreateInfo(tags request.CreateTagsInfoRequest)
	FindAllInfo() []response.TagsInfoResponse
}
