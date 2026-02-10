package logviewer

import (
	"fmt"
)

func Show(level string, tail int) error {
	path, err := LogFilePath()
	if err != nil {
		return err
	}

	lines, err := ReadAllLines(path)
	if err != nil {
		return fmt.Errorf("cannot read log file: %w", err)
	}

	lines = FilterByLevel(lines, level)
	lines = Tail(lines, tail)

	if len(lines) == 0 {
		fmt.Println("No logs found.")
		return nil
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}
