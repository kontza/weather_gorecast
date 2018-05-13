package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/spf13/viper"
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

func readConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/weather_gorecast/")
	viper.AddConfigPath("$HOME/.weather_gorecast")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func buildUrl(lat float32, lon float32) (string, error) {
	req, err := http.NewRequest("GET", "https://api.openweathermap.org/data/2.5/forecast", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("appid", viper.GetString("api_key"))
	q.Add("lat", fmt.Sprintf("%.6f", lat))
	q.Add("lon", fmt.Sprintf("%.6f", lon))
	req.URL.RawQuery = q.Encode()
	return req.URL.String(), nil
}

func (forecastService) GetForecast(lat float32, lon float32) ([]Forecast, error) {
	var err error
	if err = readConfig(); err != nil {
		return nil, err
	}
	logger := log.NewLogfmtLogger(os.Stderr)
	logger.Log("lat", lat, "lon", lon)
	var url string
	if url, err = buildUrl(lat, lon); err != nil {
		return nil, err
	}
	var response *http.Response
	if response, err = http.Get(url); err != nil {
		return nil, err
	}

	defer response.Body.Close()
	var contents []byte
	if contents, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}
	logger.Log("contents", contents)
	retVal := []Forecast{{
		1525952133, 20, 20, 20, 1002, 1002, 96, 200, 0, 7, 180, 0, 0,
	}}
	return retVal, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
