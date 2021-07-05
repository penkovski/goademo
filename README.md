# Goa Demo Service

HTTP service created with the [Goa v3](https://goa.design/) framework.
It can be used as an example and a starting point
for writing microservices with [Goa](https://goa.design/).

## main.go

The example provides a stripped down `main.go` file for wiring up
the Goa components with the service implementation.

It also shows how to use multiple *services* in the Goa DSL to
provide separate logical implementations of components inside the
micro-service. For example, `liveness` and `readiness` endpoints are
implemented in a separate component and are not mixed with the 
**weather** service business logic. The Swagger documentation is also
exposed as a separate service. 

Example:

```go
var _ = Service("weather", func() {
	Description("The weather service returns information for a given location.")
	Method("WeatherQuery", func() {
		HTTP(func() {
			POST("/v1/weather")
			...
		})
	})
})

var _ = Service("healthcheck", func() {
	Description("The healthcheck service is used report on the liveness and readiness status of the weather service.")

	Method("liveness", func() {
		HTTP(func() {
			GET("/liveness")
			...
		})
	})

	Method("readiness", func() {
		HTTP(func() {
			GET("/readiness")
			...
		})
	})
})
```

## goagen.sh

Goa code generation doesn't work when you use a vendor directory.
This means that if you explicitly build your project with vendor
support enabled, `goa gen` will fail.

The `goagen.sh` script is a work-around that limitation. 
It alters GOFLAGS so that the goa code generation doesn't look for
dependencies in the `vendor` directory.

> If you don't use the vendor mode, then you may 
> not need to use this script.

## Docker environment

You can start the service with docker-compose for easier 
local development. The dockerized environment uses a Go watcher
which automatically rebuilds the project after some `.go` files 
are changed, which may be handy.

To start the service and see the logs of the container, just type:
```shell
docker-compose up -d
docker-compose logs -f
```

**IMPORTANT:** If you want the service to run correctly, you should
provide an API key for https://openweathermap.org/api. The sample 
one `deadbeef` placed inside the `docker-composes.yml` won't work, 
as it is not a real API key provisioned by the [openweather](https://home.openweathermap.org/users/sign_up) service.

## Test the service

After the service is up and running, say for example at port 8080,
you can hit it with curl:
```shell
curl -vvv localhost:8080/v1/liveness
```