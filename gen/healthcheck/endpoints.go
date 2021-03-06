// Code generated by goa v3.3.1, DO NOT EDIT.
//
// healthcheck endpoints
//
// Command:
// $ goa gen github.com/penkovski/goademo/design

package healthcheck

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "healthcheck" service endpoints.
type Endpoints struct {
	Liveness  goa.Endpoint
	Readiness goa.Endpoint
}

// NewEndpoints wraps the methods of the "healthcheck" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Liveness:  NewLivenessEndpoint(s),
		Readiness: NewReadinessEndpoint(s),
	}
}

// Use applies the given middleware to all the "healthcheck" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Liveness = m(e.Liveness)
	e.Readiness = m(e.Readiness)
}

// NewLivenessEndpoint returns an endpoint function that calls the method
// "liveness" of service "healthcheck".
func NewLivenessEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Liveness(ctx)
	}
}

// NewReadinessEndpoint returns an endpoint function that calls the method
// "readiness" of service "healthcheck".
func NewReadinessEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Readiness(ctx)
	}
}
