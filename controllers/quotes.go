package controllers

import (
	"net/http"
	"scrape/ginger/models"

	"github.com/gin-gonic/gin"
)

// GET /quotes
// Get random quotes
func FindQuotes(c *gin.Context) {
	var quote models.QUOTE

	if err := models.DB.Order("RANDOM()").First(&quote).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": quote})
}
