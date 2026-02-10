package models

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}