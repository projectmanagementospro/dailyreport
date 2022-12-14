package config

import (
	"dailyreport/helper"
	"dailyreport/models/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	dbURL := "postgres://root:root@172.26.1.3:5432/dailyreport?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	helper.PanicIfError(err)
	db.AutoMigrate(&domain.DailyReport{}, &domain.Reports{})
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	helper.PanicIfError(err)
	dbSQL.Close()
}
