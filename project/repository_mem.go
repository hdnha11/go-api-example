package project

import (
	"strings"

	"github.com/hdnha11/go-api-example/entity"
)

// InMemRepository in memory repository
type InMemRepository struct {
	data map[string]*entity.Project
}

// NewInMemRepository create new repository
func NewInMemRepository() *InMemRepository {
	return &InMemRepository{
		data: make(map[string]*entity.Project),
	}
}

// Find find a project by id
func (i *InMemRepository) Find(id entity.ID) (*entity.Project, error) {
	p := i.data[id.String()]

	if p == nil {
		return nil, entity.ErrNotFound
	}
	return p, nil
}

// FindByName find projects by name (contains)
func (i *InMemRepository) FindByName(name string) ([]*entity.Project, error) {
	var ps []*entity.Project

	for _, p := range i.data {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(name)) {
			ps = append(ps, p)
		}
	}

	if len(ps) == 0 {
		return nil, entity.ErrNotFound
	}

	return ps, nil
}

// FindAll find all projects
func (i *InMemRepository) FindAll() ([]*entity.Project, error) {
	ps := []*entity.Project{}

	for _, p := range i.data {
		ps = append(ps, p)
	}

	return ps, nil
}

// Save create new project
func (i *InMemRepository) Save(project *entity.Project) (entity.ID, error) {
	i.data[project.ID.String()] = project

	return project.ID, nil
}

// Update project
func (i *InMemRepository) Update(id entity.ID, project *entity.Project) error {
	p, err := i.Find(id)

	if err != nil {
		return err
	}

	p.Name = project.Name
	p.Description = project.Description

	return nil
}

// Delete project
func (i *InMemRepository) Delete(id entity.ID) error {
	if _, err := i.Find(id); err != nil {
		return err
	}

	delete(i.data, id.String())

	return nil
}
