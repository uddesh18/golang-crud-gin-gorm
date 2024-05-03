package repository

import (
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

// initilazi database -> model
type TagsInfoRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements TagsInfoRepository.
func (t *TagsInfoRepositoryImpl) Delete(tagsId int) {
	panic("unimplemented")
}

// FindAll implements TagsInfoRepository.
func (t *TagsInfoRepositoryImpl) FindAllInfo() []model.TagsInfo {
	var tagsInfo []model.TagsInfo
	result := t.Db.Find(&tagsInfo)
	helper.ErrorPanic(result.Error)
	return tagsInfo
}

// FindById implements TagsInfoRepository.
func (t *TagsInfoRepositoryImpl) FindById(tagsId int) (tags model.TagsInfo, err error) {
	panic("unimplemented")
}

// Save implements TagsInfoRepository.
func (t *TagsInfoRepositoryImpl) Save(tags model.TagsInfo) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

// Update implements TagsInfoRepository.
func (t *TagsInfoRepositoryImpl) Update(tags model.TagsInfo) {
	panic("unimplemented")
}

func NewTagsInfoRepositoryImpl(Db *gorm.DB) TagsInfoRepository {
	return &TagsInfoRepositoryImpl{Db: Db}
}
