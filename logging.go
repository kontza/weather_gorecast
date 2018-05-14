package main

import (
	"time"

	"github.com/go-kit/kit/log"
	"fmt"
)

type loggingMiddleware struct {
	logger log.Logger
	next   ForecastService
}

func (mw loggingMiddleware) GetForecast(lat float32, lon float32) (output *ForecastAPIResponse, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetForecast",
			"lat", lat,
			"lon", lon,
			"output", fmt.Sprintf("%v, %d rows", output.City.Coord, output.Cnt),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.GetForecast(lat, lon)
	return
}
