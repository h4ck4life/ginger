package services

import (
	"context"

	"scrape/ginger/models"
)

func GetOneRandomQuote(ctx context.Context, quote *models.QUOTE) error {
	err := models.DB.WithContext(ctx).Order("RANDOM()").First(&quote).Error
	return err
}
