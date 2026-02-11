package cmd

import (
	"github.com/jonalphabert/db-guard/internal/doctor"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system requirements and dependencies",
	Long:  `The doctor command checks if the required database backup tools (pg_dump, mysqldump) are installed and accessible in the system PATH.`,
	Run: func(cmd *cobra.Command, args []string) {
		runDoctor()
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

func runDoctor() {
	doctor.PrintHeader()

	pgCheck := doctor.CheckExecutable("pg_dump")
	if pgCheck.Exists {
		doctor.PrintSuccess("pg_dump", pgCheck.Version)
	} else {
		pgHints := doctor.InstallHints{
			Windows: "Download PostgreSQL installer and select \"Command Line Tools\"\nhttps://www.postgresql.org/download/windows/",
			MacOS:   "brew install postgresql",
			Ubuntu:  "sudo apt install postgresql-client",
		}
		doctor.PrintFailure("pg_dump", "Required for PostgreSQL backups.", pgHints)
	}

	mysqlCheck := doctor.CheckExecutable("mysqldump")
	if mysqlCheck.Exists {
		doctor.PrintSuccess("mysqldump", mysqlCheck.Version)
	} else {
		mysqlHints := doctor.InstallHints{
			Windows: "Install MySQL Community Server and select \"Client Tools\"\nhttps://dev.mysql.com/downloads/installer/",
			MacOS:   "brew install mysql",
			Ubuntu:  "sudo apt install mysql-client",
		}
		doctor.PrintFailure("mysqldump", "Required for MySQL backups.", mysqlHints)
	}

	doctor.PrintFooter()
}
