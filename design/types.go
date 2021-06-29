package design

import . "goa.design/goa/v3/dsl"

var WeatherQueryRequest = Type("WeatherQueryRequest", func() {
	Field(1, "lat", Float64, "Latitude coordinate of location.", func() {
		Example(51.50853)
	})
	Field(2, "lon", Float64, "Longitude coordinate of location.", func() {
		Example(-0.12574)
	})
	Field(3, "units", String, "Units of measurement. Possible values are 'standard', 'metric' or 'imperial'.", func() {
		Example("metric")
	})
	Required("lat", "lon")
})

var WeatherQueryResult = Type("WeatherQueryResult", func() {
	Field(1, "temp", Float64, func() {
		Example(23.4)
	})
	Field(2, "feels_like", Float64, func() {
		Example(22.1)
	})
	Field(3, "pressure", Int, func() {
		Example(1024)
	})
	Field(4, "wind_speed", Int, func() {
		Example(10)
	})
	Required("temp")
})
