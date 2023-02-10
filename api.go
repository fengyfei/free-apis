package freeapis

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIQuery interface {
	Path() string
	Method() string
}

func Execute(o APIQuery) ([]byte, error) {
	req, err := http.NewRequest(o.Method(), o.Path(), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get %s with response code %d", o.Path(), resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
