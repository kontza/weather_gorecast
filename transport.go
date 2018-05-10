package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeUppercaseEndpoint(svc ForecastService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getForecastRequest)
		v, err := svc.GetForecast(req.lat, req.lon)
		if err != nil {
			return getForecastResponse{v, err.Error()}, nil
		}
		return getForecastResponse{v, ""}, nil
	}
}

func decodeGetForecastRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getForecastRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type getForecastRequest struct {
	lat float32 `json:"lat"`
	lon float32 `json:"lon"`
}

type getForecastResponse struct {
	V   []Forecast `json:"v"`
	Err string     `json:"err,omitempty"`
}
