package main

import (
	"fmt"
	"github.com/enekofb/metrics/pkg/config"
	"github.com/enekofb/metrics/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	defectsLastMonth = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "defects_last_month",
		Help: "defects last month",
	})

	logger = log.Default()
)

// TODO: needs to be a map
func recordMetrics(queryFuncs []metrics.QueryFunc) {
	go func() {
		for {

			for _, queryFunc := range queryFuncs {
				numDefects, err := queryFunc()
				if err != nil {
					panic(err)
				}
				defectsLastMonth.Set(float64(numDefects))
			}

			time.Sleep(2 * time.Second)
		}
	}()
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	logger.Print("reading configuration")
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic(fmt.Errorf("cannot find config path"))
	}
	config, err := config.Read(configPath)
	if err != nil {
		panic(err)
	}

	queryFuncs := metrics.CreateMetricsFromConfig(config.GithubConfig.Queries)

	recordMetrics(queryFuncs)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/healthz", http.HandlerFunc(healthz))

	http.ListenAndServe(":8080", nil)
}
