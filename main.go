package main

import (
	"net/http"
	"os"

	"time"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/patrickmn/go-cache"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "forecast_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "forecast_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "forecast_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here
	var c *cache.Cache
	c = cache.New(time.Minute, 2*time.Minute)

	var svc ForecastService
	svc = forecastService{}
	svc = cachingMiddleware{c, logger, svc}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	forecastHandler := httptransport.NewServer(
		makeForecastEndpoint(svc),
		decodeGetForecastRequest,
		encodeResponse,
	)

	http.Handle("/forecast", forecastHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":58080")
	logger.Log("err", http.ListenAndServe(":58080", nil))
}
