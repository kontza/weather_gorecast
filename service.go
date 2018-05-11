package main

import (
	"errors"
)

// ForecastService provides operations on strings.
type ForecastService interface {
	GetForecast(float32, float32) ([]Forecast, error)
}

/*
{
	"dt":1406106000,
	"main":{
		"temp":298.77,
		"temp_min":298.77,
		"temp_max":298.774,
		"pressure":1005.93,
		"sea_level":1018.18,
		"grnd_level":1005.93,
		"humidity":87,
		"temp_kf":0.26},
	"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],
	"clouds":{"all":88},
	"wind":{"speed":5.71,"deg":229.501},
	"sys":{"pod":"d"},
	"dt_txt":"2014-07-23 09:00:00"
}
*/
type Forecast struct {
	Timestamp      int32   `json:"timestamp"`
	Temperature    float32 `json:"temperature"`
	TemperatureMin float32 `json:"temperatureMin"`
	TemperatureMax float32 `json:"temperatureMax"`
	PressureSea    float32 `json:"pressureSea"`
	PressureGround float32 `json:"pressureGround"`
	Humidity       float32 `json:"humidity"`
	WeatherID      int32   `json:"weatherId"`
	Cloudiness     float32 `json:"cloudiness"`
	WindSpeed      float32 `json:"windSpeed"`
	WindDirection  float32 `json:"windDirection"`
	Rain3h         float32 `json:"rain3h"`
	Snow3h         float32 `json:"snow3h"`
}

type forecastService struct{}

func (forecastService) GetForecast(lat float32, lon float32) ([]Forecast, error) {
	retVal := []Forecast{{
		1525952133, 20, 20, 20, 1002, 1002, 96, 200, 0, 7, 180, 0, 0,
	}}
	return retVal, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
