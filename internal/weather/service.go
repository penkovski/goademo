package weather

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/penkovski/goademo/gen/weather"
)

type OpenWeather interface {
	WeatherQuery(ctx context.Context, lat, lon float64, units string) (*WeatherInfo, error)
}

type WeatherInfo struct {
	Temp      float64
	FeelsLike float64
	Pressure  int
	Humidity  int
}

type Service struct {
	logger      zerolog.Logger
	openweather OpenWeather
}

func New(logger zerolog.Logger, openWeather OpenWeather) *Service {
	return &Service{logger: logger, openweather: openWeather}
}

// WeatherQuery returns weather info for a given location.
func (s *Service) WeatherQuery(ctx context.Context, req *weather.WeatherQueryRequest) (res *weather.WeatherQueryResult, err error) {
	units := "metric"
	if req.Units != nil {
		units = *req.Units
	}

	weatherInfo, err := s.openweather.WeatherQuery(ctx, req.Lat, req.Lon, units)
	if err != nil {
		s.logger.Err(err).Msg("failed to get weather info")
		return nil, fmt.Errorf("failed to get weather info: %w", err)
	}

	return &weather.WeatherQueryResult{
		Temp:      weatherInfo.Temp,
		FeelsLike: ptrFloat64(weatherInfo.FeelsLike),
		Pressure:  ptrInt(weatherInfo.Pressure),
	}, nil
}

func ptrFloat64(f float64) *float64 {
	return &f
}

func ptrInt(i int) *int {
	return &i
}
