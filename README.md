# Introduction
This is the Go-part of the KIKY 2018 -demonstration.

# Usage
Add this to your main `docker-compose.yml`:

```yaml
  weather-gorecast:
    container_name: weather-gorecast
    build:
      context: weather-gorecast
      args:
        - PACKAGE_SOURCE=github.com/kontza/weather_gorecast
        - PACKAGE_APP=weather_gorecast
    volumes:
      - ./gorecast_config.yaml:/app/config.yaml
    ports:
      - "58080:58080"
```