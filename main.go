package main

import (
	"github.com/enekofb/metrics/pkg/config"
	"github.com/enekofb/metrics/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var (
	defectsLastMonth = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "defects_last_month",
		Help: "defects last month",
	})
)

func recordMetrics() {
	go func() {
		for {
			numDefects, err := metrics.GetLastMonthDefectMetricsByTeam()
			if err != nil {
				panic(err)
			}
			defectsLastMonth.Set(float64(numDefects))
			time.Sleep(2 * time.Second)
		}
	}()
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	err := config.Read("resources")
	if err != nil {
		panic(err)
	}
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/healthz", http.HandlerFunc(healthz))

	http.ListenAndServe(":8080", nil)
}
