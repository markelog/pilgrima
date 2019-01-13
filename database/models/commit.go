package models

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Commit model struct
type Commit struct {
	gorm.Model
	Hash     string   `gorm:"unique;not null;" json:"hash,omitempty"`
	Author   string   `json:"author,omitempty"`
	Message  string   `json:"message,omitempty"`
	Service  string   `json:"service,omitempty"`
	BranchID uint     `json:"-"`
	Reports  []Report `json:"reports,omitempty"`
}

var commitSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"hash": {"type": "string", "minLength": 1},
		"author": {"type": "string", "minLength": 1},
		"message": {"type": "string", "minLength": 1},
		"service": {"type": "string", "minLength": 1},
		"branch": {"type": "number", "minimum": 1},
		"reports": {
			"type": "array", 
			"items": {
				"type": "object"
			}
		}
	},
	"required": ["hash", "author", "message", "branch"]
}`)

// Validate model
func (commit Commit) Validate(db *gorm.DB) {
	commitLoader := gojsonschema.NewGoLoader(commit)
	check, _ := gojsonschema.Validate(commitSchema, commitLoader)

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}
}
