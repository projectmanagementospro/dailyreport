package repository

import (
	"dailyreport/models/domain"
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type DailyReportRepositoryMock struct {
	Mock mock.Mock
}

func (repository *DailyReportRepositoryMock) FindById(id uint) (domain.DailyReport, error) {
	arguments := repository.Mock.Called(id)
	var dailyreport domain.DailyReport
	if arguments.Get(0) == nil {
		return dailyreport, errors.New("id not found")
	} else {
		//dreport := arguments.Get(0).(domain.DailyReport)
		return dailyreport, nil
	}
}

func (repository *DailyReportRepositoryMock) All() []domain.DailyReport {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	} else {
		//dreport := arguments.Get(0).([]domain.DailyReport)
		var dreport []domain.DailyReport
		return dreport
	}
}

func (repository *DailyReportRepositoryMock) Create(d domain.DailyReport) domain.DailyReport {
	arguments := repository.Mock.Called()
	if arguments.Get(0) != nil {
		return d
	}
	return d
}

func (repository *DailyReportRepositoryMock) Update(d domain.DailyReport) domain.DailyReport {
	arguments := repository.Mock.Called()
	if arguments.Get(0) != nil {
		return d
	}
	return d
}

func (repository *DailyReportRepositoryMock) Delete(d domain.DailyReport) {
	arguments := repository.Mock.Called()
	if arguments.Get(0) != nil {
		fmt.Println("Deleted")
	}
}
