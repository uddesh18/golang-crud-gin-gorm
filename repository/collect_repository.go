package repository

import "golang-crud-gin/model"

type CombineStruct struct {
	Tags     model.Tags     `json:"tags"`
	TagsInfo model.TagsInfo `json:"tagsinfo"`
}

type CollectRepository interface {
	GetCombinedData(tags model.Tags, tagsinfo model.TagsInfo) CombineStruct
}
