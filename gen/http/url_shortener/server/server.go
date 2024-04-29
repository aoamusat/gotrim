// Code generated by goa v3.16.1, DO NOT EDIT.
//
// UrlShortener HTTP server
//
// Command:
// $ goa gen olayml.xyz/gotrim/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	urlshortener "olayml.xyz/gotrim/gen/url_shortener"
)

// Server lists the UrlShortener service endpoint HTTP handlers.
type Server struct {
	Mounts             []*MountPoint
	CreateShortURL     http.Handler
	GenHTTPOpenapiJSON http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the UrlShortener service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *urlshortener.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemGenHTTPOpenapiJSON http.FileSystem,
) *Server {
	if fileSystemGenHTTPOpenapiJSON == nil {
		fileSystemGenHTTPOpenapiJSON = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"CreateShortURL", "POST", "/api/v1/shorten"},
			{"./gen/http/openapi.json", "GET", "/openapi.json"},
		},
		CreateShortURL:     NewCreateShortURLHandler(e.CreateShortURL, mux, decoder, encoder, errhandler, formatter),
		GenHTTPOpenapiJSON: http.FileServer(fileSystemGenHTTPOpenapiJSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "UrlShortener" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.CreateShortURL = m(s.CreateShortURL)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return urlshortener.MethodNames[:] }

// Mount configures the mux to serve the UrlShortener endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCreateShortURLHandler(mux, h.CreateShortURL)
	MountGenHTTPOpenapiJSON(mux, goahttp.Replace("", "/./gen/http/openapi.json", h.GenHTTPOpenapiJSON))
}

// Mount configures the mux to serve the UrlShortener endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountCreateShortURLHandler configures the mux to serve the "UrlShortener"
// service "CreateShortUrl" endpoint.
func MountCreateShortURLHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/v1/shorten", f)
}

// NewCreateShortURLHandler creates a HTTP handler which loads the HTTP request
// and calls the "UrlShortener" service "CreateShortUrl" endpoint.
func NewCreateShortURLHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateShortURLRequest(mux, decoder)
		encodeResponse = EncodeCreateShortURLResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "CreateShortUrl")
		ctx = context.WithValue(ctx, goa.ServiceKey, "UrlShortener")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/openapi.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}