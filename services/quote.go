package services

import "scrape/ginger/models"

func GetOneRandomQuote(quote *models.QUOTE) error {
	err := models.DB.Order("RANDOM()").First(&quote).Error
	return err
}
