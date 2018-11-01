package models

import (
	"github.com/jinzhu/gorm"
)

type Branch struct {
	gorm.Model
	Name      string
	ProjectID uint
	Commits   []Commit
}
