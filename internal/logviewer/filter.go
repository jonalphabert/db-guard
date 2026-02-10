package logviewer

import (
	"strings"
)

func extractLevel(line string) string {
	if len(line) < 3 || line[0] != '[' {
		return ""
	}

	end := strings.Index(line, "]")
	if end == -1 {
		return ""
	}

	return line[1:end]
}

func FilterByLevel(lines []string, level string) []string {
	var result []string
	levelPriority := map[string]int{
		"DEBUG": 1,
		"INFO":  2,
		"WARN":  3,
		"ERROR": 4,
	}

	if _, ok := levelPriority[level]; !ok {
		return lines
	}

	minPriority := levelPriority[level]
	
	for _, line := range lines {
		lineLevel := extractLevel(line)
		priority, ok := levelPriority[lineLevel]
		if !ok {
			continue
		}

		if priority >= minPriority {
			result = append(result, line)
		}
	}

	return result
}
