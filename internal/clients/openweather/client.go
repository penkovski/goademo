package openweather

import (
	"context"
	"fmt"

	"github.com/penkovski/goademo/internal/weather"
)

type Client struct {
	addr   string
	apiKey string
}

func New(addr string, apiKey string) *Client {
	return &Client{
		addr:   addr,
		apiKey: apiKey,
	}
}

func (c *Client) WeatherQuery(ctx context.Context, lat, lon float64, units string) (*weather.WeatherInfo, error) {
	return nil, fmt.Errorf("not implemented")
}
