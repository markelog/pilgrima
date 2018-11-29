package models

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Branch model
type Branch struct {
	gorm.Model
	Name      string   `json:"name,omitempty"`
	ProjectID uint     `json:"project,omitempty"`
	Commits   []Commit `json:"commits,omitempty"`
}

var branchSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"name": {"type": "string", "minLength": 1},
		"project": {"type": "number", "minimum": 1},
		"commits": {
			"type": "array", 
			"items": {
				"type": "number"
			}
		}
	},
	"required": ["name", "project"]
}`)

// Validate model
func (branch Branch) Validate(db *gorm.DB) {
	branchLoader := gojsonschema.NewGoLoader(branch)
	check, _ := gojsonschema.Validate(branchSchema, branchLoader)

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}
}
