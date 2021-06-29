package main

type Config struct {
	HTTP httpConfig
}

type httpConfig struct {
	Host string `envconfig:"HTTP_HOST"`
	Port string `envconfig:"HTTP_PORT" default:"8080"`
}
