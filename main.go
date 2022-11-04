package main

import (
	"fmt"
	"github.com/enekofb/metrics/pkg/config"
	"github.com/enekofb/metrics/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"time"
)

var (
	defectsLastMonth = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "defects_last_month",
		Help: "defects last month",
	})
)

func recordMetrics(config config.MetricsConfig) {
	go func() {
		for {
			numDefects, err := metrics.GetLastMonthDefectMetricsByTeam(config)
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
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic(fmt.Errorf("cannot find config path"))
	}
	config, err := config.Read(configPath)
	if err != nil {
		panic(err)
	}
	recordMetrics(config)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/healthz", http.HandlerFunc(healthz))

	http.ListenAndServe(":8080", nil)
}
