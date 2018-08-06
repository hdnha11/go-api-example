package project

import (
	"github.com/hdnha11/go-api-example/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoRepository MongoDB repository
type MongoRepository struct {
	db *mgo.Collection
}

// NewMongoRepository create new MongoDB repository
func NewMongoRepository(db *mgo.Database) *MongoRepository {
	return &MongoRepository{db: db.C("projects")}
}

// Find find a project by id
func (m *MongoRepository) Find(id entity.ID) (*entity.Project, error) {
	p := &entity.Project{}

	if err := m.db.FindId(id).One(p); err != nil {
		return nil, err
	}

	return p, nil
}

// FindByName find projects by name (contains)
func (m *MongoRepository) FindByName(name string) ([]*entity.Project, error) {
	ps := []*entity.Project{}

	if err := m.db.Find(bson.M{"name": bson.RegEx{Pattern: name, Options: "i"}}).All(&ps); err != nil {
		return nil, err
	}

	return ps, nil
}

// FindAll find all projects
func (m *MongoRepository) FindAll() ([]*entity.Project, error) {
	ps := []*entity.Project{}

	if err := m.db.Find(bson.M{}).All(&ps); err != nil {
		return nil, err
	}

	return ps, nil
}

// Save create new project
func (m *MongoRepository) Save(project *entity.Project) (entity.ID, error) {
	if err := m.db.Insert(project); err != nil {
		return "", err
	}

	return project.ID, nil
}

// Update project
func (m *MongoRepository) Update(id entity.ID, project *entity.Project) error {
	return m.db.Update(bson.M{"_id": id}, project)
}

// Delete project
func (m *MongoRepository) Delete(id entity.ID) error {
	return m.db.Remove(bson.M{"_id": id})
}
