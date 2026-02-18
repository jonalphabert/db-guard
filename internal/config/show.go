package config

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"

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