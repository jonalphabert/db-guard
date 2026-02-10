package models

type BackupConfig struct {
	Dir       string `yaml:"dir"`
	Retention int    `yaml:"retention"`
}