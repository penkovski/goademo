package openweather

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/penkovski/goademo/internal/weather"
)

type Client struct {
	addr       string
	apiKey     string
	httpClient *http.Client
}

func New(addr string, apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		addr:       addr,
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) WeatherQuery(ctx context.Context, lat, lon float64, units string) (*weather.WeatherInfo, error) {
	url, err := url.Parse(c.addr)
	if err != nil {
		return nil, fmt.Errorf("invalid address: %v", err)
	}

	if err := c.validateUnits(units); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("lat", strconv.FormatFloat(lat, 'f', 4, 64))
	q.Add("lon", strconv.FormatFloat(lon, 'f', 4, 64))
	q.Add("units", units)
	q.Add("appid", c.apiKey)
	req.URL.RawQuery = q.Encode()

	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		return nil, newUnexpectedResponse(resp)
	}

	var response oneCallResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &weather.WeatherInfo{
		Temp:      response.Current.Temp,
		FeelsLike: response.Current.FeelsLike,
		Pressure:  response.Current.Pressure,
		Humidity:  response.Current.Humidity,
	}, nil
}

func (c *Client) validateUnits(units string) error {
	if units == "metric" || units == "standard" || units == "imperial" {
		return nil
	}
	return errors.New("invalid units")
}

type oneCallResponse struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Current struct {
		Temp      float64 `json:"temp"`       // temperature
		FeelsLike float64 `json:"feels_like"` // human perception of temperature
		Pressure  int     `json:"pressure"`   // atmospheric pressure on the sea level, hPa
		Humidity  int     `json:"humidity"`   // humidity %
	} `json:"current"`
}

func newUnexpectedResponse(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &unexpectedResponseError{
			Code: resp.StatusCode,
		}
	}
	return &unexpectedResponseError{
		Code: resp.StatusCode,
		Body: body,
	}
}

// unexpectedResponseError indicates that the remote returned an unexpected
// HTTP response code.
type unexpectedResponseError struct {
	// Code is the returned status code.
	Code int
	// Body is the content of the returned body.
	Body []byte
	// Message is the message of the error. It is populated only for structured
	// errors.
	Message string
}

// Error returns human-friendly message describing the error.
func (e *unexpectedResponseError) Error() string {
	return fmt.Sprintf("unexpected response code %d with body: %s", e.Code, e.Body)
}

// Temporary reports whether the error is temporary.
func (e *unexpectedResponseError) Temporary() bool {
	switch e.Code {
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable:
		return true
	default:
		return false
	}
}
