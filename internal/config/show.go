package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"

	"github.com/jonalphabert/db-guard/internal/logger"
	"github.com/jonalphabert/db-guard/internal/models"
	"github.com/jonalphabert/db-guard/internal/setup"
)

func showConfig(keyword string, value interface{}) {
	color.New(color.FgHiBlue).Printf("%s\t\t: ", keyword)
	color.New(color.FgWhite).Printf("%v\n", value)
}

func showSection(section string) {
	color.New(color.FgHiGreen).Printf("%s\n", section)
	color.New(color.FgHiGreen).Printf("===============================\n")
}

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

func ShowAsJson(configuration models.Config) {
	jsonConfig, err := json.MarshalIndent(configuration, "", "  ")
	if err != nil {
		logger.Error("Error marshaling config to JSON: %v", err)
		return
	}
	fmt.Println(string(jsonConfig))
}

func showBeautify(configuration models.Config) {
	showSection("Database Config")
	showConfig("Database Host", configuration.DatabaseConfig.Host)
	showConfig("Database Port", configuration.DatabaseConfig.Port)
	showConfig("Username", configuration.DatabaseConfig.User)
	showConfig("Password", configuration.DatabaseConfig.Password)
	showConfig("Database Name", configuration.DatabaseConfig.DbName)
	
	fmt.Println()
	showSection("Backup Config")
	showConfig("Backup Path", configuration.BackupConfig.Dir)
	showConfig("Retention", configuration.BackupConfig.Retention)
}

func Show(showAsJson bool, showPassword bool) {
	configPath, err := setup.ConfigPath()
	if err != nil {
		logger.Error("Error getting config path: %v", err)
		return
	}

	configuration, err := readConfig(configPath)
	if err != nil {
		logger.Error("Error reading config file: %v", err)
		return
	}
	
	if !showPassword {
		configuration.DatabaseConfig.Password = "(secret_key)"
	}
		
	if showAsJson {
		ShowAsJson(configuration)
		return
	} 

	showBeautify(configuration)
	return
}