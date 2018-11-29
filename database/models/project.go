package models

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Project model
type Project struct {
	gorm.Model
	Name       string `gorm:"not null;"`
	Repository string `gorm:"unique; not null;" json:"repository"`
	Token      *Token
	Branches   []Branch
	Users      []User `gorm:"many2many:project_users;"`
}

var projectSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"repository": {"type": "string", "minLength": 1}
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
