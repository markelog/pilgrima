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
				Hash    string `json:"hash"`
				Author  string `json:"author"`
				Message string `json:"message"`
				Report  map[string]struct {
					Size uint `json:"size"`
					Gzip uint `json:"gzip"`
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
			BranchID: branch.ID,
			Hash:     args.Project.Branch.Commit.Hash,
			Author:   args.Project.Branch.Commit.Author,
			Message:  args.Project.Branch.Commit.Message,
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
	for key, data := range args.Project.Branch.Commit.Report {
		reports = append(reports, &models.Report{
			Name: key,
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

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

// LastArgs are arguments for Last handler
type LastArgs struct {
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
}

type lastValue struct {
	Size uint `json:"size"`
	Gzip uint `json:"gzip"`
}

// LastResult is a return value for Last handler
type LastResult map[string]lastValue

// Last will get you last report
func (report *Report) Last(args *LastArgs) (result LastResult, err error) {
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

	err = report.db.Select("DISTINCT(name), size, gzip").Where(
		"commit_id = (?)",
		commit,
	).Find(&reports).Error

	if err != nil {
		return nil, err
	}

	result = make(map[string]lastValue)
	for _, report := range reports {
		result[report.Name] = lastValue{
			Size: report.Size,
			Gzip: report.Gzip,
		}
	}

	return result, err
}

// GetArgs are arguments for get handler
type GetArgs struct {
	Repository string `json:"repository"`
	Branch     string `json:"branch"`
}

type getValue struct {
	Size uint `json:"size"`
	Gzip uint `json:"gzip"`
}

// GetResult is a return value for Get handler
type GetResult map[string][]getValue

// Get reports
func (report *Report) Get(args *GetArgs) (result GetResult, err error) {
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
		).QueryExpr()
	)

	err = report.db.Select("name, size, gzip").Where(
		"commit_id in (?)",
		commit,
	).Order("created_at DESC").Find(&reports).Error

	if err != nil {
		return nil, err
	}

	result = make(map[string][]getValue)

	for _, report := range reports {
		if _, ok := result[report.Name]; !ok {
			result[report.Name] = []getValue{}
		}

		result[report.Name] = append(result[report.Name], getValue{
			Size: report.Size,
			Gzip: report.Gzip,
		})
	}

	return result, nil
}
