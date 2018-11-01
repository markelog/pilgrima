package models

import (
	"github.com/jinzhu/gorm"
)

type Commit struct {
	gorm.Model
	Committer string
	Message   string
	Pull      string
	ProjectID uint
}
