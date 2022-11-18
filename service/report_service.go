package service

import "projects/repository"

type ReportService struct {
	repo repository.Report
}

func NewReportService(repo repository.Report) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReport(year int, month int) (string, error) {
	return s.repo.GetReport(year, month)
}
