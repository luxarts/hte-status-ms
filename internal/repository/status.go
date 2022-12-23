package repository

import (
	"hte-status-ms/internal/domain"
	"log"
)

type StatusRepository interface {
	Update(p *domain.Status) error
}

type statusRepository struct {
}

func NewStatusRepository() StatusRepository {
	return &statusRepository{}
}

func (r *statusRepository) Update(p *domain.Status) error {
	log.Println(p)
	return nil
}
