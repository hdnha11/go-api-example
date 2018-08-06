package project

import (
	"github.com/hdnha11/go-api-example/entity"
)

// Repository project repository inteface
type Repository interface {
	Find(id entity.ID) (*entity.Project, error)
	FindByName(name string) ([]*entity.Project, error)
	FindAll() ([]*entity.Project, error)
	Save(project *entity.Project) (entity.ID, error)
	Update(id entity.ID, project *entity.Project) error
	Delete(id entity.ID) error
}
