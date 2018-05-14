package main

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/patrickmn/go-cache"
)

type cachingMiddleware struct {
	cache  *cache.Cache
	logger log.Logger
	next   ForecastService
}

func (mw cachingMiddleware) GetForecast(lat float32, lon float32) (output *ForecastAPIResponse, err error) {
	key := fmt.Sprintf("%.6f_%6f", lat, lon)
	if x, found := mw.cache.Get(key); found {
		mw.logger.Log("Serving from cache.")
		retVal := x.(ForecastAPIResponse)
		return &retVal, nil
	}
	output, err = mw.next.GetForecast(lat, lon)
	mw.cache.Set(key, *output, cache.DefaultExpiration)
	return
}
