package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           ForecastService
}

func (mw instrumentingMiddleware) GetForecast(lat float32, lon float32) (output []Forecast, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetForecast", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.GetForecast(lat, lon)
	return
}
