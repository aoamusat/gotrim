package gotrimapi

import (
	"context"
	"log"

	"olayml.xyz/gotrim/database/models"
	urlshortener "olayml.xyz/gotrim/gen/url_shortener"
	"olayml.xyz/gotrim/utils"
)

// UrlShortener service example implementation.
// The example methods log the requests and return zero values.
type urlShortenersrvc struct {
	logger *log.Logger
}

// NewURLShortener returns the UrlShortener service implementation.
func NewURLShortener(logger *log.Logger) urlshortener.Service {
	return &urlShortenersrvc{logger}
}

// CreateShortURL implements CreateShortUrl.
func (s *urlShortenersrvc) CreateShortURL(ctx context.Context, payload *urlshortener.CreateShortURLPayload) (res *urlshortener.Create, err error) {
	DbClient := models.GetDatabaseClient()

	DbClient.AutoMigrate(&models.User{}, &models.URL{})

	ShortUrl := utils.GenerateShortUrl(payload.LongURL)
	res = &urlshortener.Create{LongURL: &payload.LongURL, ShortURL: &ShortUrl}
	CreatedUrl := DbClient.Create(&models.URL{LongUrl: *res.LongURL, ShortUrl: *res.ShortURL, UserID: 1})
	s.logger.Printf("urlShortener.CreateShortUrl: %v", CreatedUrl)
	return res, nil
}
