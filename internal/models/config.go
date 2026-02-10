package models

type Config struct {
	DatabaseConfig DatabaseConfig `yaml:"database"`
	BackupConfig   BackupConfig   `yaml:"backup"`
}
