package main

import (
	"golang-crud-gin/config"
	"golang-crud-gin/controller"
	_ "golang-crud-gin/docs"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	"golang-crud-gin/router"
	"golang-crud-gin/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("tags_info").AutoMigrate(&model.TagsInfo{})

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)
	tagsInfoRepository := repository.NewTagsInfoRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	tagsInfoService := service.NewTagsInfoServiceImpl(tagsInfoRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)
	tagsInfoController := controller.NewTagsInfoController(tagsInfoService)

	// Router
	routes := router.NewRouter(tagsController, tagsInfoController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
