package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type GithubConfig struct {
	Token     string `yaml:"token"`
	BugLabel  string `yaml:"bug_label"`
	Owner     string `yaml:"owner"`
	Repo      string `yaml:"repo"`
	TeamLabel string `yaml:"team_label"`
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
