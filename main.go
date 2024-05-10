package main

import (
	"golang-crud-gin/config"
	"golang-crud-gin/controller"
	_ "golang-crud-gin/docs"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	"golang-crud-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
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
	db.Table("user_login_details").AutoMigrate(&model.UserLoginDetails{})
	db.Table("user_registration").AutoMigrate(&model.UserRegistration{})

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)
	tagsInfoRepository := repository.NewTagsInfoRepositoryImpl(db)
	userlogindetailRepository := repository.NewUserLoginDetailsRepositoryImpl(db)
	userregistrationRepository := repository.NewUserRegistrationImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	tagsInfoService := service.NewTagsInfoServiceImpl(tagsInfoRepository, validate)
	userlogindetailsService := service.NewUserLoginDetailServiceImpl(userlogindetailRepository, validate)
	userRegistrationService := service.NewUserRegistrationServiceImpl(userregistrationRepository, validate)

	// Controller
	tagsInfoController := controller.NewTagsInfoController(tagsInfoService)
	userlogindetailController := controller.NewLoginDetailsController(userlogindetailsService)
	userRegistrationController := controller.NewUserRegistrationController(userRegistrationService)

	// Router
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	routes := TagsRouter("/api", router)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
