package models

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Report model struct
type Report struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Size     uint   `json:"size,omitempty"`
	Gzip     uint   `json:"gzip,omitempty"`
	CommitID uint   `json:"commit,omitempty"`
}

var reportSchema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"name": {"type": "string", "minLength": 1},
		"size": {"type": "number",  "minimum": 1},
		"gzip": {"type": "number", "minimum": 1},
		"commit": {"type": "number","minimum": 1}
	},
	"required": ["name", "size", "gzip", "commit"]
}`)

// Validate model
func (report Report) Validate(db *gorm.DB) {
	reportLoader := gojsonschema.NewGoLoader(report)
	check, _ := gojsonschema.Validate(reportSchema, reportLoader)

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}
}
