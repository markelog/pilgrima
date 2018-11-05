package models

import (
	"github.com/jinzhu/gorm"
)

// Report model struct
type Report struct {
	gorm.Model
	Name   string
	Size   uint
	Commit Commit
}
