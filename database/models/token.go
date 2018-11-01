package models

import (
	"github.com/jinzhu/gorm"
)

// Token model
type Token struct {
	gorm.Model
	Token     string
	Project   Project
	ProjectID string
}
