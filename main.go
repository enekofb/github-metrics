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
	logger = log.Default()
)

type Metric struct {
	Name             string
	MetricFunc       metrics.QueryFunc
	PrometheusMetric prometheus.Gauge
}

func recordMetrics(metrics []Metric) {
	go func() {
		for {

			for _, metric := range metrics {
				logger.Println("execute %v", metric.Name)
				metricValue, err := metric.MetricFunc()
				if err != nil {
					logger.Println("error %v", err.Error())
				}
				metric.PrometheusMetric.Set(float64(metricValue))
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
	metrics := createMetrics(queryFuncs)
	recordMetrics(metrics)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/healthz", http.HandlerFunc(healthz))

	http.ListenAndServe(":8080", nil)
}

// TODO move me to a better place
func createMetrics(funcs map[string]metrics.QueryFunc) []Metric {
	var metrics []Metric
	for queryName, queryFunc := range funcs {
		metric := Metric{
			Name:       queryName,
			MetricFunc: queryFunc,
			PrometheusMetric: promauto.NewGauge(prometheus.GaugeOpts{
				Name: queryName,
				Help: queryName,
			}),
		}
		metrics = append(metrics, metric)
	}

	return metrics
}
