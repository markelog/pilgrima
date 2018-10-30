package models

import (
	"github.com/jinzhu/gorm"
)

type Report struct {
	gorm.Model
	Name     string
	Size     uint
	Service  string
	CommitID uint
}
