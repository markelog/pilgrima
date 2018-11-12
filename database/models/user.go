package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Name     string    `gorm:"not null;"`
	Projects []Project `gorm:"many2many:user_projects;"`
}
