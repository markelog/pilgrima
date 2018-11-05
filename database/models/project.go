package models

import (
	"github.com/jinzhu/gorm"
)

// Project model
type Project struct {
	gorm.Model
	Name       string `gorm:"not null;"`
	Repository string `gorm:"unique; not null;"`
	Token      *Token
	Branches   []Branch
	Users      []User
}
