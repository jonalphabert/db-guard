package config

import (
	"os"

	"github.com/jonalphabert/db-guard/internal/models"
	"gopkg.in/yaml.v3"
)

func readConfig(configPath string) (models.Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return models.Config{}, err
	}
	defer file.Close()

	var config models.Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return models.Config{}, err
	}

	return config, nil
}