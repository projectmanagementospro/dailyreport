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

func (drs *dailyreportService) All() []domain.DailyReport {
	return drs.dailyreportRepository.All()
}

func (drs *dailyreportService) Create(request web.DailyReportRequest) (domain.DailyReport, error) {
	dailyreport := domain.DailyReport{}
	err := smapping.FillStruct(&dailyreport, smapping.MapFields(&request))

	if err != nil {
		return dailyreport, err
	}

	// _, err = drs.dailyreportRepository.IsDuplicateEmail(request.Email)
	// if err != nil {
	// 	return dailyreport, err
	// }
	return drs.dailyreportRepository.Create(dailyreport), nil
}

func (drs *dailyreportService) Update(request web.DailyReportUpdateRequest) (domain.DailyReport, error) {
	dailyreport := domain.DailyReport{}
	res, err := drs.dailyreportRepository.FindById(request.ID)
	if err != nil {
		return dailyreport, err
	}
	err = smapping.FillStruct(&dailyreport, smapping.MapFields(&request))
	if err != nil {
		return dailyreport, err
	}
	dailyreport.ReportNumber = res.ReportNumber
	dailyreport.User_id = res.User_id
	return drs.dailyreportRepository.Update(dailyreport), nil
}

func (drs *dailyreportService) FindById(id uint) (domain.DailyReport, error) {
	dailyreport, err := drs.dailyreportRepository.FindById(id)
	if err != nil {
		return dailyreport, err
	}
	return dailyreport, nil
}

func (drs *dailyreportService) Delete(id uint) error {
	dailyreport, err := drs.dailyreportRepository.FindById(id)
	if err != nil {
		return err
	}
	drs.dailyreportRepository.Delete(dailyreport)
	return nil
}
