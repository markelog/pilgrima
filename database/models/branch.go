package models

import (
	"github.com/jinzhu/gorm"
)

// Branch model
type Branch struct {
	gorm.Model
	Name      string
	Commits   []Commit
	ProjectID uint
}
