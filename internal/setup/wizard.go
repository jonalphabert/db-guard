package setup

import "github.com/AlecAivazis/survey/v2"

type SetupInput struct {
	DBType    string
	Host      string
	Port      int
	DBName    string
	User      string
	Password  string
	Retention int
	BackupDir string
}

func RunWizard() (*SetupInput, error) {
	input := &SetupInput{}

	qs := []*survey.Question{
		{
			Name: "DBType",
			Prompt: &survey.Select{
				Message: "Database type:",
				Options: []string{"postgres", "mysql"},
				Default: "postgres",
			},
		},
		{
			Name: "Host",
			Prompt: &survey.Input{
				Message: "Host:",
				Default: "localhost",
			},
		},
		{
			Name: "Port",
			Prompt: &survey.Input{
				Message: "Port:",
				Default: "5432",
			},
		},
		{
			Name: "DBName",
			Prompt: &survey.Input{
				Message: "Database name:",
			},
		},
		{
			Name: "User",
			Prompt: &survey.Input{
				Message: "Username:",
			},
		},
		{
			Name: "Password",
			Prompt: &survey.Password{
				Message: "Password:",
			},
		},
		{
			Name: "Retention",
			Prompt: &survey.Input{
				Message: "Backup retention (days):",
				Default: "3",
			},
		},
		{
			Name: "BackupDir",
			Prompt: &survey.Input{
				Message: "Backup directory:",
				Default: "./backups",
			},
		},
	}

	err := survey.Ask(qs, input)
	return input, err
}
