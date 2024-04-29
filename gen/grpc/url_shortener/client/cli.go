// Code generated by goa v3.16.1, DO NOT EDIT.
//
// UrlShortener gRPC client CLI support package
//
// Command:
// $ goa gen olayml.xyz/gotrim/design

package client

import (
	"encoding/json"
	"fmt"

	url_shortenerpb "olayml.xyz/gotrim/gen/grpc/url_shortener/pb"
	urlshortener "olayml.xyz/gotrim/gen/url_shortener"
)

// BuildCreateShortURLPayload builds the payload for the UrlShortener
// CreateShortUrl endpoint from CLI flags.
func BuildCreateShortURLPayload(urlShortenerCreateShortURLMessage string) (*urlshortener.CreateShortURLPayload, error) {
	var err error
	var message url_shortenerpb.CreateShortURLRequest
	{
		if urlShortenerCreateShortURLMessage != "" {
			err = json.Unmarshal([]byte(urlShortenerCreateShortURLMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"long_url\": \"Quia architecto odio odio qui.\"\n   }'")
			}
		}
	}
	v := &urlshortener.CreateShortURLPayload{
		LongURL: message.LongUrl,
	}

	return v, nil
}
