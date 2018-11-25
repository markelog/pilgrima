package report

import (
	"errors"

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
func (report *Report) Create(args *CreateArgs) (err error) {
	var (
		project models.Project
		branch  models.Branch
		commit  = &models.Commit{
			BranchID:  branch.ID,
			Hash:      args.Project.Branch.Commit.Hash,
			Committer: args.Project.Branch.Commit.Committer,
			Message:   args.Project.Branch.Commit.Message,
		}

		tx = report.db.Begin()
	)

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
		Hash:     args.Project.Branch.Commit.Hash,
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

	spew.Dump(args.Project.Branch.Commit.Report)

	if len(reports) == 0 {
		tx.Rollback()
		return errors.New("There is no applicable reports")
	}

	err = tx.Model(&commit).Association("Reports").Append(reports).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

// LastArgs are arguments to last get report in the branch
type LastArgs struct {
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
}

// LastResult return value of Last
type LastResult struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Gzip int    `json:"gzip"`
}

// Last will get you last report
func (report *Report) Last(args *LastArgs) (result []LastResult, err error) {
	var (
		reports []models.Report

		project = report.db.Table("projects").Select("id").Where(
			"repository = ?",
			args.Repository,
		).QueryExpr()

		branch = report.db.Table("branches").Select("id").Where(
			"name = ? AND project_id = (?)",
			args.Branch, project,
		).QueryExpr()

		commit = report.db.Table("commits").Select("id").Where(
			"branch_id = (?)",
			branch,
		).Order("created_at DESC").Limit(1).QueryExpr()
	)

	err = report.db.Select("DISTINCT(name), size, gzip").Where("commit_id = (?)", commit).
		Find(&reports).Error

	for _, report := range reports {
		result = append(result, LastResult{
			Name: report.Name,
			Size: report.Size,
			Gzip: report.Gzip,
		})
	}

	return
}
