// Code generated by goa v3.16.1, DO NOT EDIT.
//
// UrlShortener gRPC server
//
// Command:
// $ goa gen olayml.xyz/gotrim/design

package server

import (
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	url_shortenerpb "olayml.xyz/gotrim/gen/grpc/url_shortener/pb"
	urlshortener "olayml.xyz/gotrim/gen/url_shortener"
)

// Server implements the url_shortenerpb.URLShortenerServer interface.
type Server struct {
	CreateShortURLH goagrpc.UnaryHandler
	url_shortenerpb.UnimplementedURLShortenerServer
}

// New instantiates the server struct with the UrlShortener service endpoints.
func New(e *urlshortener.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		CreateShortURLH: NewCreateShortURLHandler(e.CreateShortURL, uh),
	}
}

// NewCreateShortURLHandler creates a gRPC handler which serves the
// "UrlShortener" service "CreateShortUrl" endpoint.
func NewCreateShortURLHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateShortURLRequest, EncodeCreateShortURLResponse)
	}
	return h
}

// CreateShortURL implements the "CreateShortURL" method in
// url_shortenerpb.URLShortenerServer interface.
func (s *Server) CreateShortURL(ctx context.Context, message *url_shortenerpb.CreateShortURLRequest) (*url_shortenerpb.CreateShortURLResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "CreateShortUrl")
	ctx = context.WithValue(ctx, goa.ServiceKey, "UrlShortener")
	resp, err := s.CreateShortURLH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*url_shortenerpb.CreateShortURLResponse), nil
}
