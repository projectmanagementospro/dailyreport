package service

import (
	"dailyreport/models/domain"
	"dailyreport/models/web"
	"dailyreport/repository"

	"github.com/mashingan/smapping"
)

type ReportsService interface {
	All() []domain.Reports
	Create(b web.ReportsRequest) (domain.Reports, error)
	FindById(id uint) (domain.Reports, error)
	Update(b web.ReportsUpdateRequest) (domain.Reports, error)
	Delete(id uint) error
}

type reportsService struct {
	reportsRepository repository.ReportsRepository
}

func NewReportsService(reportsRepository repository.ReportsRepository) ReportsService {
	return &reportsService{reportsRepository: reportsRepository}
}

func (s *reportsService) All() []domain.Reports {
	return s.reportsRepository.All()
}

func (s *reportsService) Create(request web.ReportsRequest) (domain.Reports, error) {
	reports := domain.Reports{}
	err := smapping.FillStruct(&reports, smapping.MapFields(&request))

	if err != nil {
		return reports, err
	}

	// _, err = s.reportsRepository.IsDuplicateEmail(request.Email)
	// if err != nil {
	// return reports, err
	// }
	return s.reportsRepository.Create(reports), nil
}

func (s *reportsService) Update(b web.ReportsUpdateRequest) (domain.Reports, error) {
	reports := domain.Reports{}
	res, err := s.reportsRepository.FindById(b.ID)
	if err != nil {
		return reports, err
	}
	err = smapping.FillStruct(&reports, smapping.MapFields(&b))
	if err != nil {
		return reports, err
	}
	reports.UserId = res.UserId
	return s.reportsRepository.Update(reports), nil
}

func (s *reportsService) FindById(id uint) (domain.Reports, error) {
	reports, err := s.reportsRepository.FindById(id)
	if err != nil {
		return reports, err
	}
	return reports, nil
}

func (s *reportsService) Delete(id uint) error {
	reports, err := s.reportsRepository.FindById(id)
	if err != nil {
		return err
	}
	s.reportsRepository.Delete(reports)
	return nil
}
