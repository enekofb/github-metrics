package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type QueryConfig struct {
	Name      string `yaml:"name"`
	Owner     string `yaml:"owner"`
	Repo      string `yaml:"repo"`
	BugLabel  string `yaml:"bug_label"`
	TeamLabel string `yaml:"team_label"`
}

type GithubConfig struct {
	Token   string        `yaml:"token"`
	Queries []QueryConfig `yaml:"queries"`
}

type MetricsConfig struct {
	GithubConfig GithubConfig `yaml:"github"`
}

func Read(configPath string) (MetricsConfig, error) {

	metricsConfig := MetricsConfig{}

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
