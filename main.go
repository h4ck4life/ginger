package main

import (
	"fmt"
	"net/http"
	"os"
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
		c.JSON(http.StatusOK, gin.H{
			"data":   "get /random quote from over 499,500 records",
			"source": "https://github.com/h4ck4life/ginger",
			"author": "@h4ck4life",
		})
	})

	r.GET("/random", controllers.FindQuotes)

	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
