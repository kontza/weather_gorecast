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
	Timestamp      int32
	Temperature    float32
	TemperatureMin float32
	TemperatureMax float32
	PressureSea    float32
	PressureGround float32
	Humidity       float32
	WeatherID      int32
	Cloudiness     float32
	WindSpeed      float32
	WindDirection  float32
	Rain3h         float32
	Snow3h         float32
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
