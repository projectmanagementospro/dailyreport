package service

import (
	"dailyreport/models/domain"
	"dailyreport/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var dreportRepository = &repository.DailyReportRepositoryMock{Mock: mock.Mock{}}
var dreportService = dailyreportService{dailyreportRepository: dreportRepository}

func TestDailyReportService_GetNotFound(t *testing.T) {

	// program mock
	dreportRepository.Mock.On("FindById", uint(1)).Return(nil)

	dailyreport, err := dreportService.FindById(1)
	assert.Nil(t, dailyreport)
	assert.NotNil(t, err)

}

func TestDailyReportService_GetSuccess(t *testing.T) {
	dailyreport := domain.DailyReport{
		Name:        "Day 07",
		Description: "No Time To Die",
		User_id:     1,
		UpdatedBy:   "Someone",
		DeletedBy:   "Unknown",
	}
	dreportRepository.Mock.On("FindById", uint(1)).Return(dailyreport, nil)

	result, err := dreportService.FindById(1)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, dailyreport.Name, result.Name)
	assert.Equal(t, dailyreport.Description, result.Description)
	assert.Equal(t, dailyreport.User_id, result.User_id)
	assert.Equal(t, dailyreport.UpdatedBy, result.UpdatedBy)
	assert.Equal(t, dailyreport.DeletedBy, result.DeletedBy)
}
