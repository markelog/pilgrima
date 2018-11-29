package models

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Project model
type Project struct {
	gorm.Model
	Name       string `gorm:"not null;" json:"name,omitempty"`
	Repository string `gorm:"unique; not null;" json:"repository,omitempty"`
	Token      *Token
	Branches   []Branch `json:"branches,omitempty"`
	Users      []User   `gorm:"many2many:project_users;"`
}

var projectSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"repository": {"type": "string", "minLength": 1},
		"name": {"type": "string", "minLength": 1},
		"branches": {
			"type": "array", 
			"items": {
				"type": "number"
			}
		},
		"users": {
			"type": "array", 
			"items": {
				"type": "number"
			}
		}
	},
	"required": ["repository"]
}`)

// Validate model
func (project Project) Validate(db *gorm.DB) {
	projectLoader := gojsonschema.NewGoLoader(project)
	check, _ := gojsonschema.Validate(projectSchema, projectLoader)

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}
}
