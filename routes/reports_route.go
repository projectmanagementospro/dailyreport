package routes

import (
	"dailyreport/injector"
	"dailyreport/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewReportsRoutes(db *gorm.DB, route *gin.Engine) {
	reportsController := injector.InitReports(db)
	reportsRoute := route.Group("/api/v1/reports")
	reportsRoute.Use(middleware.ErrorHandler())
	reportsRoute.Use(cors.Default())
	reportsRoute.GET("/", reportsController.All)
	reportsRoute.GET("/:id", reportsController.FindById)
	reportsRoute.POST("/", reportsController.Insert)
	reportsRoute.PUT("/:id", reportsController.Update)
	reportsRoute.DELETE("/:id", reportsController.Delete)
}
