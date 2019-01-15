package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ConnectArgs arguments for the Connect()
type ConnectArgs struct {
	Host, Port, Name, SSL, Database, User, Password string
}

// Connect to database
func Connect(args *ConnectArgs) (db *gorm.DB, err error) {
	credentials := fmt.Sprintf(
		"host=%s port=%s dbname=%s sslmode=%s user=%s password=%s",
		args.Host, args.Port, args.Name, args.SSL, args.User, args.Password,
	)

	db, err = gorm.Open("postgres", credentials)
	if err != nil {
		return
	}

	return
}

// ConnectDSN connect to database through DSN string
func ConnectDSN(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", dsn)

	if err != nil {
		return
	}

	return
}
