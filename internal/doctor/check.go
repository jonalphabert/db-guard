package doctor

import (
	"context"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

type CheckResult struct {
	Name    string
	Exists  bool
	Version string
}

func CheckExecutable(name string) CheckResult {
	var result CheckResult
	result.Name = name

	path, err := exec.LookPath(name)

	if err != nil {
		result.Exists = false
		return result
	}

	if path == "" {
		result.Exists = false
		return result
	}

	result.Exists = true
	result.Version = getVersion(name)

	return result
}

func getVersion(name string) string {
	var output []byte
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, "--version")
	output, err = cmd.Output()

	if err != nil {
		return "unknown"
	}

	return parseVersion(name, string(output))
}

func parseVersion(name string, output string) string {
	var version string

	if strings.Contains(name, "pg_dump") {
		version = parsePostgresVersion(output)
	} else if strings.Contains(name, "mysqldump") {
		version = parseMySQLVersion(output)
	} else {
		version = parseGenericVersion(output)
	}

	return version
}

func parsePostgresVersion(output string) string {
	re := regexp.MustCompile(`\(PostgreSQL\) (\d+(\.\d+)*)`)

	matches := re.FindStringSubmatch(output)

	if len(matches) > 1 {
		return matches[1]
	}

	reSimple := regexp.MustCompile(`pg_dump.*\s(\d+(\.\d+)*)`)
	matchesSimple := reSimple.FindStringSubmatch(output)
	if len(matchesSimple) > 1 {
		return matchesSimple[1]
	}

	return "unknown"
}

func parseMySQLVersion(output string) string {
	re := regexp.MustCompile(`Ver (\d+(\.\d+)*)`)

	matches := re.FindStringSubmatch(output)

	if len(matches) > 1 {
		return matches[1]
	}

	return "unknown"
}

func parseGenericVersion(output string) string {
	re := regexp.MustCompile(`(\d+(\.\d+)+)`)

	matches := re.FindStringSubmatch(output)

	if len(matches) > 1 {
		return matches[1]
	}

	return strings.TrimSpace(output)
}
