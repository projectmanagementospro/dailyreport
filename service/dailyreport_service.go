package service

import (
	"dailyreport/models/domain"
	"dailyreport/models/web"
	"dailyreport/repository"

	"github.com/mashingan/smapping"
)

type DailyReportService interface {
	All() []domain.DailyReport
	Create(b web.DailyReportRequest) (domain.DailyReport, error)
	FindById(id uint) (domain.DailyReport, error)
	Update(b web.DailyReportUpdateRequest) (domain.DailyReport, error)
	Delete(id uint) error
}

type dailyreportService struct {
	dailyreportRepository repository.DailyReportRepository
}

func NewDailyReportService(dailyreportRepository repository.DailyReportRepository) DailyReportService {
	return &dailyreportService{dailyreportRepository: dailyreportRepository}
}

func (s *dailyreportService) All() []domain.DailyReport {
	return s.dailyreportRepository.All()
}

func (s *dailyreportService) Create(request web.DailyReportRequest) (domain.DailyReport, error) {
	dailyreport := domain.DailyReport{}
	err := smapping.FillStruct(&dailyreport, smapping.MapFields(&request))

	if err != nil {
		return dailyreport, err
	}

	// _, err = s.dailyreportRepository.IsDuplicateEmail(request.Email)
	// if err != nil {
	// 	return dailyreport, err
	// }
	return s.dailyreportRepository.Create(dailyreport), nil
}

func (s *dailyreportService) Update(b web.DailyReportUpdateRequest) (domain.DailyReport, error) {
	dailyreport := domain.DailyReport{}
	res, err := s.dailyreportRepository.FindById(b.ID)
	if err != nil {
		return dailyreport, err
	}
	err = smapping.FillStruct(&dailyreport, smapping.MapFields(&b))
	if err != nil {
		return dailyreport, err
	}
	//dailyreport.ID = res.ID
	dailyreport.User_id = res.User_id
	return s.dailyreportRepository.Update(dailyreport), nil
}

func (s *dailyreportService) FindById(id uint) (domain.DailyReport, error) {
	dailyreport, err := s.dailyreportRepository.FindById(id)
	if err != nil {
		return dailyreport, err
	}
	return dailyreport, nil
}

func (s *dailyreportService) Delete(id uint) error {
	dailyreport, err := s.dailyreportRepository.FindById(id)
	if err != nil {
		return err
	}
	s.dailyreportRepository.Delete(dailyreport)
	return nil
}
