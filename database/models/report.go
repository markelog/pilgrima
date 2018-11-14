package models

import (
	"time"

	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/xeipuuv/gojsonschema"
)

// Report model struct
type Report struct {
	gorm.Model
	Name      string `json:"name"`
	Size      int    `json:"size"`
	Gzip      int    `json:"gzip"`
	CommitID  int
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

var schema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"name": {"type": "string", "minLength": 1},
		"size": {"type": "number", "minLength": 1},
		"gzip": {"type": "number", "minLength": 1}
	},
	"required": ["name", "size", "gzip"]
}`)

// Validate model
func (report Report) Validate(db *gorm.DB) {
	reportLoader := gojsonschema.NewGoLoader(report)
	check, _ := gojsonschema.Validate(schema, reportLoader)

	for _, desc := range check.Errors() {
		db.AddError(errors.New(desc.String()))
	}
}
