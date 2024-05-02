package service

import "golang-crud-gin/data/request"

type TagsInfoService interface {
	CreateInfo(tags request.CreateTagsInfoRequest)
}
