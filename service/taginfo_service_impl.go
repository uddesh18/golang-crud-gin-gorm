package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type TagsInfoServiceImpl struct {
	TagsInfoRepository repository.TagsInfoRepository
	Validate           *validator.Validate
}

// CreateInfo implements TagsInfoService.
func (t *TagsInfoServiceImpl) CreateInfo(tags request.CreateTagsInfoRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.TagsInfo{
		Tid:    tags.Tid,
		Class:  tags.Class,
		Bg:     tags.Bg,
		Dob:    tags.Dob,
		Gender: tags.Gender,
	}
	t.TagsInfoRepository.Save(tagModel)
}

func NewTagsInfoServiceImpl(tagInfoRepository repository.TagsInfoRepository, validate *validator.Validate) TagsInfoService {
	return &TagsInfoServiceImpl{
		TagsInfoRepository: tagInfoRepository,
		Validate:           validate,
	}
}
