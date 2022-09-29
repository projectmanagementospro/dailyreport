package routes

import (
	"dailyreport/injector"
	"dailyreport/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewDailyReportRoutes(db *gorm.DB, route *gin.Engine) {
	dreportController := injector.InitDailyReport(db)
	dreportRoute := route.Group("/api/v1/dreport")
	dreportRoute.Use(middleware.ErrorHandler())
	dreportRoute.Use(cors.Default())
	dreportRoute.GET("/", dreportController.All)
	dreportRoute.GET("/:id", dreportController.FindById)
	dreportRoute.POST("/", dreportController.Insert)
	dreportRoute.PUT("/:id", dreportController.Update)
	dreportRoute.DELETE("/:id", dreportController.Delete)
}
