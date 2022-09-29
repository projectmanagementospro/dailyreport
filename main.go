package main

import (
	"dailyreport/config"
	"dailyreport/helper"
	"dailyreport/middleware"
	"dailyreport/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.NewDB()
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	router := NewRouter()
	log.Fatal(router.Run(":" + os.Getenv("GO_PORT")))
}

func NewRouter() *gin.Engine {
	err := godotenv.Load()
	helper.PanicIfError(err)
	/**
	@description Setup Database Connection
	*/

	/**
	@description Init Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/
	if os.Getenv("GO_ENV") != "production" && os.Getenv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if os.Getenv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/**
	@description Setup Middleware
	*/

	/**
	@description Init All Route
	*/
	routes.NewDailyReportRoutes(db, router)
	router.Use(middleware.ErrorHandler())
	router.Use(cors.Default())

	return router
}
