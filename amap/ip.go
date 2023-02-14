package amap

import (
	"fmt"
	"net/http"
)

type IpLocRequest struct {
	IP string // Required
}

func (o *IpLocRequest) Path() string {
	return fmt.Sprintf("https://restapi.amap.com/v3/ip?key=%s&ip=%s", Key(), o.IP)
}

func (o *IpLocRequest) Method() string {
	return http.MethodGet
}
