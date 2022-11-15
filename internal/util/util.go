package util

import (
	"net/http"
	"time"
)

var hc = &http.Client{Timeout: 300 * time.Millisecond}

func Client() *http.Client {
	return hc
}
