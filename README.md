# Introduction
This is the Go-part of the KIKY 2018 -demonstration.

# Usage
Add this to your main `docker-compose.yml`:

```yaml
  gorecast:
    container_name: gorecast
    build:
      context: weather_gorecast
      args:
        - PACKAGE_SOURCE=github.com/kontza/weather_gorecast
        - PACKAGE_APP=weather_gorecast
```