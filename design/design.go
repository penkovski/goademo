package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("weather", func() {
	Title("Weather Service")
	Description("Service for weather information.")
	Server("weather", func() {
		Description("Weather service")
		Host("staging", func() {
			Description("Staging server")
			URI("https://dev.myweatherdomain.com")
		})
		Host("development", func() {
			Description("Local development server")
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("weather", func() {
	Description("The weather service returns information for a given location.")
	Method("WeatherQuery", func() {
		Payload(WeatherQueryRequest)
		Result(WeatherQueryResult)
		HTTP(func() {
			POST("/v1/weather")
			Response(StatusOK)
			Response(StatusBadRequest)
			Response(StatusInternalServerError)
		})
	})
})

var _ = Service("healthcheck", func() {
	Description("The healthcheck service is used report on the liveness and readiness status of the weather service.")

	Method("liveness", func() {
		Result(func() {
			Field(1, "status", String, func() {
				Example("up")
			})
			Required("status")
		})

		HTTP(func() {
			GET("/liveness")
			Response(StatusOK)
		})
	})

	Method("readiness", func() {
		Result(func() {
			Field(1, "status", String, func() {
				Example("up")
			})
			Required("status")
		})

		Error("ServiceUnavailableError", func() {
			Description("ServiceUnavailableError is returned when the service is unavailable.")
			Field(1, "status", String, func() {
				Example("status")
			})
			Required("status")
		})

		HTTP(func() {
			GET("/readiness")
			Response(StatusOK)
			Response("ServiceUnavailableError", StatusServiceUnavailable)
		})
	})
})

var _ = Service("openapi", func() {
	Description("The openapi service serves the OpenAPI(v3) definition.")
	Meta("swagger:generate", "false")
	HTTP(func() {
		Path("/swagger-ui")
	})
	Files("/openapi.json", "./gen/http/openapi3.json", func() {
		Description("JSON document containing the OpenAPI(v3) service definition")
	})
	Files("/{*filepath}", "./swagger/")
})
