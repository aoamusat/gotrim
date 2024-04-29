package design

import (
	"fmt"

	. "goa.design/goa/v3/dsl"
)

const (
	HTTP_PORT = 3050
	GRPC_PORT = 3060
)

type URLResult struct {
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

var ShortenedURL = ResultType("ShortenedURL", func() {
	Description("Shortened URL response")
	Attribute("long_url", String, "Original Long URL", func() {
		Example("https://example.com/very/long/url/path")
	})
	Attribute("short_url", String, "Shortened URL", func() {
		Example("https://short.url/xyz123")
	})
	Required("short_url", "long_url")
})

var _ = API("GoTrim API", func() {
	Title("GoTrim API Service")
	Description("URL shortening service written in Go")
	Server("gotrim", func() {
		Host("localhost", func() {
			URI(fmt.Sprintf("http://localhost:%v", HTTP_PORT))
			URI(fmt.Sprintf("grpc://localhost:%v", GRPC_PORT))
		})
	})
})

var _ = Service("create", func() {
	Description("The create service performs shortening operation on long URL.")

	Method("create", func() {
		Payload(func() {
			Field(1, "long_url", String, "Payload for shortening a URL")
			Required("long_url")
		})

		Result(ShortenedURL)

		HTTP(func() {
			POST("/api/v1/shorten")
			Body(func() {
				Attribute("long_url", String, "URL to shorten")
				Required("long_url")
			})
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
