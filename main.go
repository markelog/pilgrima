package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/markelog/pilgrima/log"
	"github.com/markelog/pilgrima/models"
)

var project models.Project

func main() {
	log := log.Log()

	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	var (
		port    = os.Getenv("PORT")
		address = ":" + port
	)

	db, err := models.Connect(
		&models.ConnectArgs{
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			SSL:      os.Getenv("DATABASE_SSL"),
		},
	)
	if err != nil {
		log.Panic(err)
	}

	db.Create(&models.Project{Name: "test"})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db.First(&project)

		fmt.Fprintf(w, project.Name)
	})

	log.WithFields(logrus.Fields{
		"port": port,
	}).Info("Started")

	log.Fatal(http.ListenAndServe(address, nil))
}
