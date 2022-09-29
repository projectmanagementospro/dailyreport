//go:build wireinject
// +build wireinject

package injector

import (
	"dailyreport/controller"
	"dailyreport/repository"
	"dailyreport/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var dailyreportSet = wire.NewSet(
	repository.NewDailyReportRepository,
	service.NewDailyReportService,
	controller.NewDailyReportController,
)

func InitDailyReport(db *gorm.DB) controller.DailyReportController {
	wire.Build(
		dailyreportSet,
	)
	return nil
}
