// Code generated by goa v3.16.1, DO NOT EDIT.
//
// UrlShortener client
//
// Command:
// $ goa gen olayml.xyz/gotrim/design

package urlshortener

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "UrlShortener" service client.
type Client struct {
	CreateShortURLEndpoint goa.Endpoint
}

// NewClient initializes a "UrlShortener" service client given the endpoints.
func NewClient(createShortURL goa.Endpoint) *Client {
	return &Client{
		CreateShortURLEndpoint: createShortURL,
	}
}

// CreateShortURL calls the "CreateShortUrl" endpoint of the "UrlShortener"
// service.
func (c *Client) CreateShortURL(ctx context.Context, p *CreateShortURLPayload) (res *Create, err error) {
	var ires any
	ires, err = c.CreateShortURLEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Create), nil
}
