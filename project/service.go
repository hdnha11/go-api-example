package project

import (
	"time"

	"github.com/hdnha11/go-api-example/entity"
)

// Service project servive
type Service struct {
	repo Repository
}

// NewService create new project service
func NewService(r Repository) *Service {
	return &Service{repo: r}
}

// Find find project by id
func (s *Service) Find(id entity.ID) (*entity.Project, error) {
	return s.repo.Find(id)
}

// FindByName find project by name
func (s *Service) FindByName(name string) ([]*entity.Project, error) {
	return s.repo.FindByName(name)
}

// FindAll find all projects
func (s *Service) FindAll() ([]*entity.Project, error) {
	return s.repo.FindAll()
}

// Save create new project
func (s *Service) Save(p *entity.Project) (entity.ID, error) {
	p.ID = entity.NewID()
	p.CreatedAt = time.Now()
	return s.repo.Save(p)
}

// Update project
func (s *Service) Update(id entity.ID, p *entity.Project) error {
	p.ID = id
	return s.repo.Update(id, p)
}

// Delete project
func (s *Service) Delete(id entity.ID) error {
	return s.repo.Delete(id)
}
