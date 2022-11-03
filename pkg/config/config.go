package config

import "github.com/spf13/viper"

type GithubConfig struct {
	token     string `yaml:"token"`
	bugLabel  string `yaml:"bug_label"`
	owner     string `yaml:"owner"`
	repo      string `yaml:"repo"`
	teamLabel string `yaml:"team_label"`
}

type MetricsConfig struct {
	githubConfig GithubConfig `yaml:"github"`
}

func Read(configPath string) (*MetricsConfig, error) {
	viper.SetConfigName("config.yaml") // name of config file (without extension)
	viper.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configPath)    // optionally look for config in the working directory
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {
		return nil, err
	}
	conf := &MetricsConfig{}

	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil

}
