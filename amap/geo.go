package amap

import (
	"fmt"
	"net/http"
)

type GeoCodeRequest struct {
	Address string // Required
	City    string // Optional
}

func (o *GeoCodeRequest) Path() string {
	path := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s&address=%s", Key(), o.Address)

	if o.City != "" {
		path += "&city=" + o.City
	}

	return path
}

func (o *GeoCodeRequest) Method() string {
	return http.MethodGet
}
