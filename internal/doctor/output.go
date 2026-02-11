package doctor

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type InstallHints struct {
	Windows string
	MacOS   string
	Ubuntu  string
}

func PrintSuccess(toolName string, version string) {
	green := color.New(color.FgGreen).SprintFunc()

	message := fmt.Sprintf(" %s %s found (v%s)", green("✓"), toolName, version)

	fmt.Println(message)
}

func PrintFailure(toolName string, purpose string, hints InstallHints) {
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf(" %s %s not found\n", red("✗"), toolName)
	fmt.Printf("    → %s\n", yellow(purpose))
	fmt.Printf("    → %s\n", yellow("Install:"))

	printPlatformHint("Windows", hints.Windows)
	printPlatformHint("macOS", hints.MacOS)
	printPlatformHint("Ubuntu", hints.Ubuntu)

	fmt.Println()
}

func PrintHeader() {
	fmt.Println("Running system checks...")
	fmt.Println()
}

func PrintFooter() {
	fmt.Println("Doctor finished.")
}

func printPlatformHint(platform string, hint string) {
	lines := strings.Split(hint, "\n")
	label := platform + ":"

	if len(lines) > 0 {
		fmt.Printf("       • %-9s %s\n", label, lines[0])
	}

	for i := 1; i < len(lines); i++ {
		fmt.Printf("         %s\n", lines[i])
	}
}
