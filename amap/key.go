package amap

import "os"

func Key() string {
	return os.Getenv("AMAP_CLIENT_KEY")
}
