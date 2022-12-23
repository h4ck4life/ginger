package main

import (
	"fmt"
	"net/http"
	"scrape/ginger/controllers"
	"scrape/ginger/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS middleware to allow all origins
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	})

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "go /quotes to get a random quote"})
	})

	r.GET("/quotes", controllers.FindQuotes)

	err := r.Run()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
