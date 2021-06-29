package weather

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/penkovski/goademo/gen/weather"
)

type Service struct {
	logger zerolog.Logger
}

func New(logger zerolog.Logger) *Service {
	return &Service{logger: logger}
}

// WeatherQuery returns weather info for a given location.
func (s *Service) WeatherQuery(context.Context, *weather.WeatherQueryRequest) (res *weather.WeatherQueryResult, err error) {
	return nil, nil
}
