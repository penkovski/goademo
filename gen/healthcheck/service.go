// Code generated by goa v3.4.3, DO NOT EDIT.
//
// healthcheck service
//
// Command:
// $ goa gen github.com/penkovski/goademo/design

package healthcheck

import (
	"context"
)

// The healthcheck service is used report on the liveness and readiness status
// of the weather service.
type Service interface {
	// Liveness implements liveness.
	Liveness(context.Context) (res *LivenessResult, err error)
	// Readiness implements readiness.
	Readiness(context.Context) (res *ReadinessResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "healthcheck"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"liveness", "readiness"}

// LivenessResult is the result type of the healthcheck service liveness method.
type LivenessResult struct {
	Status string
}

// ReadinessResult is the result type of the healthcheck service readiness
// method.
type ReadinessResult struct {
	Status string
}

// ServiceUnavailableError is returned when the service is unavailable.
type ServiceUnavailableError struct {
	Status string
}

// Error returns an error description.
func (e *ServiceUnavailableError) Error() string {
	return "ServiceUnavailableError is returned when the service is unavailable."
}

// ErrorName returns "ServiceUnavailableError".
func (e *ServiceUnavailableError) ErrorName() string {
	return "ServiceUnavailableError"
}