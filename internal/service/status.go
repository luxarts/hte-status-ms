package service

import (
	"hte-status-ms/internal/domain"
	"hte-status-ms/internal/repository"
)

type StatusService interface {
	Update(p *domain.Status) (*domain.Status, error)
}

type statusService struct {
	repo repository.StatusRepository
}

func NewStatusService(repo repository.StatusRepository) StatusService {
	return &statusService{repo: repo}
}

func (r *statusService) Update(p *domain.Status) (*domain.Status, error) {
	err := r.repo.Update(p)
	return p, err
}
