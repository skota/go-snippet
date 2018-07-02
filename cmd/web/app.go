package main

import "../../data/models"

//struct to hold data shared between handlers
type App struct {
	Database  *models.Database
	HTMLDir   string
	StaticDir string
}
