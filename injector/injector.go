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

var reportsSet = wire.NewSet(
	repository.NewReportsRepository,
	service.NewReportsService,
	controller.NewReportsController,
)

func InitDailyReport(db *gorm.DB) controller.DailyReportController {
	wire.Build(
		dailyreportSet,
	)
	return nil
}

func InitReports(db *gorm.DB) controller.ReportsController {
	wire.Build(
		reportsSet,
	)
	return nil
}
