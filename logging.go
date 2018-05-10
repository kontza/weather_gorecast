package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   ForecastService
}

func (mw loggingMiddleware) GetForecast(lat float32, lon float32) (output []Forecast, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetForecast",
			"lat", lat,
			"lon", lon,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.GetForecast(lat, lon)
	return
}
