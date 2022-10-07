package service

import (
	"dailyreport/models/domain"
	"dailyreport/models/web"
	"dailyreport/repository"

	"github.com/mashingan/smapping"
)

type DailyReportService interface {
	All() []domain.DailyReport
	Create(request web.DailyReportRequest) (domain.DailyReport, error)
	FindById(id uint) (domain.DailyReport, error)
	Update(request web.DailyReportUpdateRequest) (domain.DailyReport, error)
	Delete(id uint) error
}

type dailyreportService struct {
	dailyreportRepository repository.DailyReportRepository
}

func NewDailyReportService(dailyreportRepository repository.DailyReportRepository) DailyReportService {
	return &dailyreportService{dailyreportRepository: dailyreportRepository}
}

func (dReportService *dailyreportService) All() []domain.DailyReport {
	return dReportService.dailyreportRepository.All()
}

func (dReportService *dailyreportService) Create(request web.DailyReportRequest) (domain.DailyReport, error) {
	dailyreport := domain.DailyReport{}
	err := smapping.FillStruct(&dailyreport, smapping.MapFields(&request))

	if err != nil {
		return dailyreport, err
	}

	// _, err = dReportService.dailyreportRepository.IsDuplicateEmail(request.Email)
	// if err != nil {
	// 	return dailyreport, err
	// }
	return dReportService.dailyreportRepository.Create(dailyreport), nil
}

func (dReportService *dailyreportService) Update(request web.DailyReportUpdateRequest) (domain.DailyReport, error) {
	dailyreport := domain.DailyReport{}
	res, err := dReportService.dailyreportRepository.FindById(request.ID)
	if err != nil {
		return dailyreport, err
	}
	err = smapping.FillStruct(&dailyreport, smapping.MapFields(&request))
	if err != nil {
		return dailyreport, err
	}
	dailyreport.UserId = res.UserId
	return dReportService.dailyreportRepository.Update(dailyreport), nil
}

func (dReportService *dailyreportService) FindById(id uint) (domain.DailyReport, error) {
	dailyreport, err := dReportService.dailyreportRepository.FindById(id)
	if err != nil {
		return dailyreport, err
	}
	return dailyreport, nil
}

func (dReportService *dailyreportService) Delete(id uint) error {
	dailyreport, err := dReportService.dailyreportRepository.FindById(id)
	if err != nil {
		return err
	}
	dReportService.dailyreportRepository.Delete(dailyreport)
	return nil
}
