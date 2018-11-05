package project

import (
	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/controllers/token"
	"github.com/markelog/pilgrima/database/models"
)

// Project type
type Project struct {
	Name       string
	Repository string
	db         *gorm.DB
	model      *models.Project
}

// New Project
func New(name, repository string, db *gorm.DB) *Project {
	return &Project{
		Name:       name,
		Repository: repository,
		db:         db,
	}
}

// Create project
func (project *Project) Create() (*gorm.DB, *models.Project) {
	project.model = &models.Project{
		Name:       project.Name,
		Repository: project.Repository,
		Token:      token.New(project.db).Model,
	}

	result := project.db.Create(project.model)

	return result, project.model
}
