// Code generated by goa v3.3.1, DO NOT EDIT.
//
// weather HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/penkovski/goademo/design

package server

import (
	"context"
	"io"
	"net/http"

	weather "github.com/penkovski/goademo/gen/weather"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeWeatherQueryResponse returns an encoder for responses returned by the
// weather WeatherQuery endpoint.
func EncodeWeatherQueryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*weather.WeatherQueryResult)
		enc := encoder(ctx, w)
		body := NewWeatherQueryOKResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeWeatherQueryRequest returns a decoder for requests sent to the weather
// WeatherQuery endpoint.
func DecodeWeatherQueryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body WeatherQueryRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateWeatherQueryRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewWeatherQueryRequest(&body)

		return payload, nil
	}
}
