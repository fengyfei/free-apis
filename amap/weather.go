package amap

import (
	"fmt"
	"net/http"
)

type WeatherRequest struct {
	City string // Required
}

func (o *WeatherRequest) Path() string {
	return fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s&extensions=all", Key(), o.City)
}

func (o *WeatherRequest) Method() string {
	return http.MethodGet
}
