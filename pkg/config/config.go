package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type QueryConfig struct {
	MetricName string `yaml:"metric_name"`
	Owner      string `yaml:"owner"`
	Repo       string `yaml:"repo"`
	BugLabel   string `yaml:"bug_label"`
	TeamLabel  string `yaml:"team_label"`
}

type GithubConfig struct {
	Token   string        `yaml:"token"`
	Queries []QueryConfig `yaml:"queries"`
}

type MetricConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type Config struct {
	GithubConfig  GithubConfig   `yaml:"github"`
	MetricsConfig []MetricConfig `yaml:"metrics"`
}

func Read(configPath string) (Config, error) {

	metricsConfig := Config{}

	metricsConfigBytes, err := os.ReadFile(configPath)

	if err != nil {
		return metricsConfig, err
	}

	err = yaml.Unmarshal(metricsConfigBytes, &metricsConfig)
	if err != nil {
		return metricsConfig, err
	}

	return metricsConfig, nil

}
