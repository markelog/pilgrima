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
func (project *Project) Create(name, repository string) (*gorm.DB, *models.Project) {
	project.Model = &models.Project{
		Name:       name,
		Repository: repository,
		Token:      token.New(project.db).Model,
	}

	result := project.db.Create(project.Model)

	return result, project.Model
}
