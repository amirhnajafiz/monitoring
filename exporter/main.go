package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	response struct {
		Requests     int       `json:"requests"`
		Logs         int       `json:"logs"`
		ResponseTime []float64 `json:"response_time"`
	}

	metrics struct {
		Requests *prometheus.Counter
		Logs     *prometheus.Counter
		Response prometheus.Histogram
	}
)

func pull(port, interval int) error {
	// register metrics

	// make http request to metrics endpoint
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("127.0.0.1:%d", port), nil)
	if err != nil {
		return err
	}

	// update metrics
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":2112", nil); err != nil {
		panic(err)
	}
}
