package token

import (
	"crypto/rand"
	"fmt"

	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/markelog/pilgrima/database/models"
)

// Token type
type Token struct {
	Token   string
	project int
	db      *gorm.DB
	model   *models.Token
}

var (
	// ErrNoSuchProject error
	ErrNoSuchProject = errors.New("No such project")
)

func generate() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)

	return fmt.Sprintf("%x", bytes)
}

// New Token
func New(project int, db *gorm.DB) *Token {
	return &Token{
		Token:   generate(),
		project: project,
		db:      db,
	}
}

// Create token
func (token *Token) Create() (string, error) {
	var project models.Project

	token.db.Where("id = ?", token.project).First(&project)

	if project.ID == 0 {
		return "", ErrNoSuchProject
	}

	token.model = &models.Token{
		Token:   token.Token,
		Project: project,
	}

	err := token.db.Create(&token.model).Error
	if err != nil {
		return "", err
	}

	return token.Token, nil
}
