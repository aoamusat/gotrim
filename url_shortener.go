package gotrimapi

import (
	"context"
	"log"

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
func (s *urlShortenersrvc) CreateShortURL(ctx context.Context, p *urlshortener.CreateShortURLPayload) (res []byte, err error) {
	s.logger.Print("urlShortener.CreateShortUrl")
	ResponseData := utils.UrlResponse{LongUrl: p.LongURL, ShortUrl: "https://go.trim/xyz"}
	return ResponseData.ToJSON()
}
