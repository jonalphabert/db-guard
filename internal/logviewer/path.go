package logviewer

import (
	"os"
	"path/filepath"
)

func LogFilePath() (string, error)  {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	logPath := filepath.Join(home, ".dbkeeper", "logs", "dbkeeper.log")
	return logPath, nil
}
