package report

import (
	"github.com/jinzhu/gorm"
)

// Report type
type Report struct {
	db    *gorm.DB
	model *gorm.DB
}

// CreateArgs are create arguments for report type
type CreateArgs struct {
	Project struct {
		Repository string `json:"repository"`
		Branch     struct {
			Name   string `json:"name"`
			Commit struct {
				Hash      string `json:"hash"`
				Committer string `json:"committer"`
				Message   string `json:"message"`
				Report    []struct {
					Name string `json:"name"`
					Size int    `json:"size"`
				} `json:"report"`
			} `json:"commit"`
		} `json:"branch"`
	} `json:"project"`
}

// New report
func New(db *gorm.DB) *Report {
	return &Report{
		db: db,
	}
}

// Create report and associated data
func (report *Report) Create(args *CreateArgs) *gorm.DB {
	return report.db.Save(args.Project)
}
