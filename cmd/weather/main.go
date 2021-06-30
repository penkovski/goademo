package main

import (
	"net/http"
	"os"
	"time"

	"github.com/penkovski/goademo/internal/clients/openweather"

	"github.com/kelseyhightower/envconfig"
	"github.com/penkovski/graceful"
	"github.com/rs/zerolog"
	goahttp "goa.design/goa/v3/http"

	ghealthcheck "github.com/penkovski/goademo/gen/healthcheck"
	ghealthchecksrv "github.com/penkovski/goademo/gen/http/healthcheck/server"
	openapisvr "github.com/penkovski/goademo/gen/http/openapi/server"
	gweathersrv "github.com/penkovski/goademo/gen/http/weather/server"
	gweather "github.com/penkovski/goademo/gen/weather"
	"github.com/penkovski/goademo/internal/healthcheck"
	"github.com/penkovski/goademo/internal/weather"
)

func main() {
	logger := zerolog.New(os.Stderr)

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		logger.Fatal().Err(err).Send()
	}

	// create clients and service dependencies
	openweather := openweather.New(cfg.OpenWeather.Addr, cfg.OpenWeather.APIKey)

	var (
		weatherSvc     gweather.Service
		healthcheckSvc ghealthcheck.Service
	)
	{
		weatherSvc = weather.New(logger, openweather)
		healthcheckSvc = healthcheck.New()
	}

	var (
		weatherEndpoints     *gweather.Endpoints
		healthcheckEndpoints *ghealthcheck.Endpoints
	)
	{
		weatherEndpoints = gweather.NewEndpoints(weatherSvc)
		healthcheckEndpoints = ghealthcheck.NewEndpoints(healthcheckSvc)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	mux := goahttp.NewMuxer()

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		weatherServer     *gweathersrv.Server
		healthcheckServer *ghealthchecksrv.Server
	)
	{
		weatherServer = gweathersrv.New(weatherEndpoints, mux, dec, enc, nil, nil)
		healthcheckServer = ghealthchecksrv.New(healthcheckEndpoints, mux, dec, enc, nil, nil)
	}

	// Configure the mux.
	gweathersrv.Mount(mux, weatherServer)
	ghealthchecksrv.Mount(mux, healthcheckServer)
	openapisvr.Mount(mux)

	var handler http.Handler = mux

	srv := &http.Server{
		Addr:    cfg.HTTP.Host + ":" + cfg.HTTP.Port,
		Handler: handler,
	}

	for _, m := range weatherServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	logger.Info().Str("host", cfg.HTTP.Host).Str("port", cfg.HTTP.Port).Msg("starting server")

	if err := graceful.Shutdown(srv, 10*time.Second); err != nil {
		logger.Error().Err(err)
	}

	logger.Info().Msg("bye bye")
}
