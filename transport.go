package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeForecastEndpoint(svc ForecastService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Coordinates)
		v, err := svc.GetForecast(req.Lat, req.Lon)
		if err != nil {
			foo := ForecastAPIResponse{}
			return getForecastResponse{foo, err.Error()}, nil
		}
		return getForecastResponse{*v, ""}, nil
	}
}

func decodeGetForecastRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request Coordinates
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type getForecastResponse struct {
	V   ForecastAPIResponse `json:"forecast"`
	Err string              `json:"err,omitempty"`
}
