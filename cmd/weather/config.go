package main

type Config struct {
	HTTP        httpConfig
	OpenWeather openWeatherConfig
}

type httpConfig struct {
	Host string `envconfig:"HTTP_HOST"`
	Port string `envconfig:"HTTP_PORT" default:"8080"`
}

type openWeatherConfig struct {
	Addr   string `envconfig:"OPENWEATHER_ADDR" required:"true"`
	APIKey string `envconfig:"OPENWEATHER_APIKEY" required:"true"`
}
