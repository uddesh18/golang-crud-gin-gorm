package repository

import "golang-crud-gin/model"

type TagsInfoRepository interface {
	Save(tags model.TagsInfo)
	Update(tags model.TagsInfo)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.TagsInfo, err error)
	FindAllInfo() []model.TagsInfo
}
