// Code generated by goa v3.16.1, DO NOT EDIT.
//
// create endpoints
//
// Command:
// $ goa gen olayml.xyz/gotrim/design

package create

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "create" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
}

// NewEndpoints wraps the methods of the "create" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Create: NewCreateEndpoint(s),
	}
}

// Use applies the given middleware to all the "create" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "create".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreatePayload)
		return s.Create(ctx, p)
	}
}