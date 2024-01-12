package controllers

import (
	"net/http"
	"scrape/ginger/models"
	"scrape/ginger/services"

	"github.com/gin-gonic/gin"
)

// GET /quotes
// Get random quotes
func FindQuotes(c *gin.Context) {
	ctx := c.Request.Context()
	var quote models.QUOTE

	err := services.GetOneRandomQuote(ctx, &quote)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": quote})
}
