package design

import (
	"fmt"

	. "goa.design/goa/v3/dsl"
)

const (
	HTTP_PORT = 3050
	GRPC_PORT = 3060
)

var _ = API("GoTrim API", func() {
	Title("GoTrim API Service")
	Description("URL shortening service written in Go")
	Server("gotrim", func() {
		Host("development", func() {
			Description("Development hosts.")
			URI(fmt.Sprintf("http://localhost:%v", HTTP_PORT))
			URI(fmt.Sprintf("grpc://localhost:%v", GRPC_PORT))
		})

		Host("production", func() {
			Description("Production hosts.")
			// URIs can be parameterized using {param} notation.
			URI("https://api.gotrim.olayml.xyz/{version}/")
			URI("grpcs://api.gotrim.olayml.xyz/{version}/")

			// Variable describes a URI variable.
			Variable("version", String, "API version", func() {
				// URL parameters must have a default value and/or an
				// enum validation.
				Default("v1")
			})

		})
	})
})

var _ = Service("UrlShortener", func() {
	Method("CreateShortUrl", func() {
		Payload(func() {
			Field(1, "long_url", String, "Payload for shortening a URL")
			Required("long_url")
		})

		Result(Bytes)

		HTTP(func() {
			POST("/api/v1/shorten")
			Body(func() {
				Attribute("long_url", String, "URL to shorten")
				Required("long_url")
			})
			Response(StatusCreated)
		})
		GRPC(func() {
		})
	})
	Files("/openapi.json", "./gen/http/openapi.json")
})
