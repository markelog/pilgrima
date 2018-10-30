package models

import (
	"github.com/jinzhu/gorm"
)

// Project model
type Project struct {
	gorm.Model
	ID       string `gorm:"primary_key"`
	Name     string `gorm:"not null;"`
	Branches []Branch
	Users    []User
}
