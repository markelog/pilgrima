package report

import (
	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/database/models"
)

type Report struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Report {
	return &Report{
		db: db,
	}
}

func Create(project unit) {
	&models.Branch{}
}
