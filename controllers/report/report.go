package report

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/database/models"
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
					Gzip int    `json:"gzip"`
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
	var project models.Project
	var branch models.Branch
	// var commit models.Commit

	report.db.
		Where(models.Project{
			Repository: args.Project.Repository,
		}).FirstOrCreate(&project)

	report.db.Where(models.Branch{
		ProjectID: project.ID,
		Name:      args.Project.Branch.Name,
	}, project.ID).FirstOrCreate(&branch)

	t := report.db.Model(&branch).Association(
		"Commits",
	).Append(&models.Commit{
		BranchID:  branch.ID,
		Hash:      args.Project.Branch.Commit.Hash,
		Committer: args.Project.Branch.Commit.Committer,
		Message:   args.Project.Branch.Commit.Message,
	})

	// report.db.Where(models.Commit{
	// 	BranchID:  branch.ID,
	// 	Hash:      args.Project.Branch.Commit.Hash,
	// 	Committer: args.Project.Branch.Commit.Committer,
	// 	Message:   args.Project.Branch.Commit.Message,
	// }).FirstOrCreate(&commit)

	spew.Dump(t)

	return report.db
}
