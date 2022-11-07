package metrics

import (
	"github.com/enekofb/metrics/pkg/config"
	"github.com/enekofb/metrics/pkg/issues"
	github2 "github.com/google/go-github/v48/github"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
)

type QueryFunc func() (int, error)

var logger = log.Default()

type Metric struct {
	Name             string
	MetricFuncs      []QueryFunc
	PrometheusMetric prometheus.Gauge
}

func CreateMetricsFromConfig(conf config.Config) []Metric {
	metricsByMetricName := createMetrics(conf.MetricsConfig)
	queryFuncByMetricName := createQueryFuncs(conf.GithubConfig.Queries)
	var metrics []Metric
	for metricName, metric := range metricsByMetricName {
		m := Metric{
			Name:             metricName,
			MetricFuncs:      queryFuncByMetricName[metricName],
			PrometheusMetric: metric.PrometheusMetric,
		}
		metrics = append(metrics, m)
	}
	return metrics
}

// TODO move me to a better place
func createMetrics(metricsConfig []config.MetricConfig) map[string]Metric {
	var metrics map[string]Metric
	for _, metricConfig := range metricsConfig {
		metric := Metric{
			Name: metricConfig.Name,
			PrometheusMetric: promauto.NewGauge(prometheus.GaugeOpts{
				Name: metricConfig.Name,
			}),
		}
		metrics[metric.Name] = metric
	}

	return metrics
}

func createQueryFuncs(queriesConfig []config.QueryConfig) map[string][]QueryFunc {
	var queryFuncMap map[string][]QueryFunc

	logger.Print("creating metrics from configuration")
	for _, queryConfig := range queriesConfig {
		logger.Println("create query function for %v", queryConfig)
		queryFunc := createMetricFuncFromConfig(queryConfig)
		queryFuncMap[queryConfig.MetricName] = append(queryFuncMap[queryConfig.MetricName], queryFunc)
	}
	return queryFuncMap
}

func createMetricFuncFromConfig(queryConfig config.QueryConfig) func() (int, error) {

	return func() (int, error) {

		owner := queryConfig.Owner
		if owner == "" {
			return -1, nil
		}

		repo := queryConfig.Repo
		if repo == "" {
			return -1, nil
		}

		defectLabel := queryConfig.BugLabel
		if defectLabel == "" {
			return -1, nil
		}

		teamLabel := queryConfig.TeamLabel
		if teamLabel == "" {
			return -1, nil
		}

		//oneDay := time.Hour * 24
		//lastMonth := time.Now().Add(-oneDay)

		options := &github2.IssueListByRepoOptions{
			Labels: []string{
				defectLabel,
				teamLabel,
			},
			//Since: lastMonth,
		}

		repoIssues, err := issues.GetRepoIssues(owner, repo, options)
		if err != nil {
			return -1, nil
		}
		return len(repoIssues), nil

	}
}
