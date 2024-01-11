package models

import (
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("quotes_v1.sqlite"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	/*
		err = database.AutoMigrate(&Quote{})
		if err != nil {
			return
		}
	*/

	DB = database
	if err := database.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}
}
