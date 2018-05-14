package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/spf13/viper"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"os"
)

// ForecastService provides operations on strings.
type ForecastService interface {
	GetForecast(float32, float32) (*ForecastAPIResponse, error)
}

/*
{
    "city":
    {
        "id": 1851632,
        "name": "Shuzenji",
        "coord":
        {
            "lon": 138.933334,
            "lat": 34.966671
        },
        "country": "JP",
        "cod": "200",
        "message": 0.0045,
        "cnt": 38,
        "list": [
        {
            "dt": 1406106000,
            "main":
            {
                "temp": 298.77,
                "temp_min": 298.77,
                "temp_max": 298.774,
                "pressure": 1005.93,
                "sea_level": 1018.18,
                "grnd_level": 1005.93,
                "humidity": 87,
                "temp_kf": 0.26
            },
            "weather": [
            {
                "id": 804,
                "main": "Clouds",
                "description": "overcast clouds",
                "icon": "04d"
            }],
            "clouds":
            {
                "all": 88
            },
            "wind":
            {
                "speed": 5.71,
                "deg": 229.501
            },
            "sys":
            {
                "pod": "d"
            },
            "dt_txt": "2014-07-23 09:00:00"
        }]
    }
}*/
type Coordinates struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}
type mainForecast struct {
	Temp                float32 `json:"temp"`
	TempMin             float32 `json:"temp_min"`
	TempMax             float32 `json:"temp_max"`
	Pressure            float32 `json:"pressure"`
	SeaLevelPressure    float32 `json:"sea_level"`
	GroundLevelPressure float32 `json:"grnd_level"`
	Humidity            float32 `json:"humidity"`
	TempKF              float32 `json:"temp_kf"`
}
type weather struct {
	Id          int32  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type clouds struct {
	All int32 `json:"all"`
}
type wind struct {
	Speed     float32 `json:"speed"`
	Direction float32 `json:"deg"`
}
type forecast struct {
	Dt           int64        `json:"dt"`
	MainForecast mainForecast `json:"main"`
	Weather      weather      `json:"weather"`
	Clouds       clouds       `json:"clouds"`
	Wind         wind         `json:"wind"`
}
type CityForecast struct {
	Id         int32       `json:"id"`
	Name       string      `json:"name"`
	Coord      Coordinates `json:"coord"`
	Country    string      `json:"country"`
	Population int32       `json:"population"`
}
type ForecastAPIResponse struct {
	City    CityForecast `json:"city"`
	List    []forecast   `json:"list"`
	Cnt     int32        `json:"cnt"`
	Message float32      `json:"message"`
	Code    int32        `json:"cod"`
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
	q.Add("units", "metric")
	req.URL.RawQuery = q.Encode()
	return req.URL.String(), nil
}

func (forecastService) GetForecast(lat float32, lon float32) (*ForecastAPIResponse, error) {
	logger := log.NewLogfmtLogger(os.Stderr)
	var err error
	if err = readConfig(); err != nil {
		return nil, err
	}
	var url string
	if url, err = buildUrl(lat, lon); err != nil {
		return nil, err
	}
	logger.Log("lat", lat, "lon", lon, "url", url)
	var response *http.Response
	if response, err = http.Get(url); err != nil {
		return nil, err
	}

	defer response.Body.Close()
	var contents []byte
	if contents, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}
	var retVal ForecastAPIResponse
	err = json.Unmarshal(contents, &retVal)
	logger.Log("lat", lat, "lon", lon, "url", url, fmt.Sprintf("%v", retVal.City.Coord))
	return &retVal, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
