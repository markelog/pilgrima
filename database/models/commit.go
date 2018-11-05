package models

import (
	"github.com/jinzhu/gorm"
)

// Commit model struct
type Commit struct {
	gorm.Model
	Committer string
	Message   string
	Pull      string
	Service   string
	Project   Project
	Reports   []Report
}
