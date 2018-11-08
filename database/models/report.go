package models

import (
	"github.com/go-errors/errors"

	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Report model struct
type Report struct {
	gorm.Model
	Name     string `json:"name"`
	Size     uint   `json:"size"`
	CommitID uint
}

var schema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"name": {"type": "string", "minLength": 1},
		"repository": {"type": "string", "format": "uri"}
	},
	"required": ["name", "repository"]
}`)

// Validate model
func (report Report) Validate(db *gorm.DB) {
	reportLoader := gojsonschema.NewGoLoader(report)
	check, _ := gojsonschema.Validate(schema, reportLoader)

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}
}
