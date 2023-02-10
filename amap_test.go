package freeapis

import (
	"fmt"
	"testing"

	"github.com/fengyfei/free-apis/amap"
)

func TestGeoCode(t *testing.T) {
	o := []APIQuery{
		&amap.GeoCodeRequest{
			Address: "保定市华北电力大学",
		},
		&amap.GeoCodeRequest{
			Address: "华北电力大学",
			City:    "北京",
		},
	}

	for _, item := range o {
		if data, err := Execute(item); err != nil {
			t.Errorf("geocode(%v) error %s", item, err.Error())
		} else {
			fmt.Println(string(data))
		}
	}
}

func TestReGeo(t *testing.T) {
	o := []APIQuery{
		&amap.ReGeoRequest{
			Location: "115.512414,38.888687",
		},
		&amap.ReGeoRequest{
			Location: "116.310015,40.089676",
			POI:      []string{"160100"},
			Radius:   3000,
		},
	}

	for _, item := range o {
		if data, err := Execute(item); err != nil {
			t.Errorf("regeo(%v) error %s", item, err.Error())
		} else {
			fmt.Println(string(data))
		}
	}
}
