package amap

import (
	"fmt"
	"net/http"
	"strconv"
)

type ReGeoRequest struct {
	Location  string   // Required
	POI       []string // Optional
	Radius    int      // Optional
	RoadLevel int      // Optional
}

func (o *ReGeoRequest) Path() string {
	path := fmt.Sprintf("https://restapi.amap.com/v3/geocode/regeo?key=%s&location=%s", Key(), o.Location)

	if len(o.POI) > 0 {
		path += "&extensions=all&poitype=" + o.POI[0]

		for i := 1; i < len(o.POI); i++ {
			path += "|" + o.POI[i]
		}
		path += "&roadlevel" + strconv.Itoa(o.RoadLevel)
	}

	if o.Radius > 0 {
		path += "&radius=" + strconv.Itoa(o.Radius)
	}

	return path
}

func (o *ReGeoRequest) Method() string {
	return http.MethodGet
}
