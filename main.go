package main

import (
	"fmt"
	"github.com/enekofb/metrics/pkg/config"
	"github.com/enekofb/metrics/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	logger = log.Default()
)

func recordMetrics(metrics []metrics.Metric) {
	go func() {
		for {
			for _, metric := range metrics {
				logger.Println("execute %v", metric.Name)
				for _, metricFunc := range metric.MetricFuncs {
					metricValue, err := metricFunc()
					if err != nil {
						logger.Println("error %v", err.Error())
					}
					metric.PrometheusMetric.Set(float64(metricValue))
				}
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

	metrics := metrics.CreateMetricsFromConfig(config)
	recordMetrics(metrics)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/healthz", http.HandlerFunc(healthz))
	http.ListenAndServe(":8080", nil)
}
