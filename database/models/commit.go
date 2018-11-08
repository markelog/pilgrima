package models

import (
	"github.com/jinzhu/gorm"
)

// Commit model struct
type Commit struct {
	gorm.Model
	Hash      string `gorm:"unique;not null;"`
	Committer string
	Message   string
	Service   string
	BranchID  uint
	Reports   []Report
}
