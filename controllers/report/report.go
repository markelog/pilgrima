package report

import (
	"errors"

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
func (report *Report) Create(args *CreateArgs) (err error) {
	var project models.Project
	var branch models.Branch
	commit := &models.Commit{
		BranchID:  branch.ID,
		Hash:      args.Project.Branch.Commit.Hash,
		Committer: args.Project.Branch.Commit.Committer,
		Message:   args.Project.Branch.Commit.Message,
	}

	var tx = report.db.Begin()

	err = tx.Where(models.Project{
		Repository: args.Project.Repository,
	}).FirstOrCreate(&project).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where(models.Branch{
		ProjectID: project.ID,
		Name:      args.Project.Branch.Name,
	}).FirstOrCreate(&branch).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where(models.Commit{
		BranchID: branch.ID,
	}).FirstOrCreate(&commit).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	reports := []*models.Report{}
	for _, data := range args.Project.Branch.Commit.Report {
		reports = append(reports, &models.Report{
			Name: data.Name,
			Size: data.Size,
			Gzip: data.Gzip,
		})
	}

	if len(reports) == 0 {
		tx.Rollback()
		return errors.New("There is not applicable reports")
	}

	err = tx.Model(&commit).Association("Reports").Append(reports).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
