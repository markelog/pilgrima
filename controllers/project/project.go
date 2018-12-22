package project

import (
	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/controllers/token"
	"github.com/markelog/pilgrima/database/models"
)

// Project type
type Project struct {
	db    *gorm.DB
	Model *models.Project
}

// New Project
func New(db *gorm.DB) *Project {
	return &Project{
		db: db,
	}
}

// Create project
func (project *Project) Create(name, repository string) (*models.Project, error) {
	project.Model = &models.Project{
		Name:       name,
		Repository: repository,
		Token:      token.New(project.db).Model,
	}

	result := project.db.Create(project.Model)

	if result.Error != nil {
		return nil, result.Error
	}

	return project.Model, nil
}

// ListValue result value for List() method
type ListValue struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

// List projects
func (project *Project) List() ([]ListValue, error) {
	var (
		projects []models.Project
		result   []ListValue
	)

	err := project.db.Find(&projects).Error
	if err != nil {
		return nil, err
	}

	for _, project := range projects {
		result = append(result, ListValue{
			Name:       project.Name,
			Repository: project.Repository,
		})
	}

	return result, nil
}
