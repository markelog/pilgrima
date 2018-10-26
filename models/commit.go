package models

import (
	"github.com/jinzhu/gorm"
)

type Commit struct {
	gorm.Model
	Name      string
	Size      uint
	ProjectID uint
}
