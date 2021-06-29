package healthcheck

import (
	"context"

	"github.com/penkovski/goademo/gen/healthcheck"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

// Liveness implements liveness.
func (s *Service) Liveness(context.Context) (res *healthcheck.LivenessResult, err error) {
	return &healthcheck.LivenessResult{Status: "up"}, nil
}

// Readiness implements readiness.
func (s *Service) Readiness(context.Context) (res *healthcheck.ReadinessResult, err error) {
	return &healthcheck.ReadinessResult{Status: "up"}, nil
}
