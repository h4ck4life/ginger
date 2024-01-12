package models

import (
	"time"

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

	// Get generic database object sql.DB to use its functions
	sqlDB, err := database.DB()
	if err != nil {
		panic("failed to get database")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	//defer sqlDB.Close()

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
