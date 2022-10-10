package service

import (
	"dailyreport/models/domain"
	"dailyreport/models/web"
	"dailyreport/repository"

	"github.com/mashingan/smapping"
)

type ReportsService interface {
	All() []domain.Reports
	Create(req web.ReportsRequest) (domain.Reports, error)
	FindById(id uint) (domain.Reports, error)
	Update(req web.ReportsUpdateRequest) (domain.Reports, error)
	Delete(id uint) error
}

type reportsService struct {
	reportsRepository repository.ReportsRepository
}

func NewReportsService(reportsRepository repository.ReportsRepository) ReportsService {
	return &reportsService{reportsRepository: reportsRepository}
}

func (reportsService *reportsService) All() []domain.Reports {
	return reportsService.reportsRepository.All()
}

func (reportsService *reportsService) Create(request web.ReportsRequest) (domain.Reports, error) {
	reports := domain.Reports{}
	err := smapping.FillStruct(&reports, smapping.MapFields(&request))

	if err != nil {
		return reports, err
	}

	return reportsService.reportsRepository.Create(reports), nil
}

func (reportsService *reportsService) Update(req web.ReportsUpdateRequest) (domain.Reports, error) {
	reports := domain.Reports{}
	res, err := reportsService.reportsRepository.FindById(req.ID)
	if err != nil {
		return reports, err
	}
	err = smapping.FillStruct(&reports, smapping.MapFields(&req))
	if err != nil {
		return reports, err
	}
	reports.UserId = res.UserId
	return reportsService.reportsRepository.Update(reports), nil
}

func (reportsService *reportsService) FindById(id uint) (domain.Reports, error) {
	reports, err := reportsService.reportsRepository.FindById(id)
	if err != nil {
		return reports, err
	}
	return reports, nil
}

func (reportsService *reportsService) Delete(id uint) error {
	reports, err := reportsService.reportsRepository.FindById(id)
	if err != nil {
		return err
	}
	reportsService.reportsRepository.Delete(reports)
	return nil
}
