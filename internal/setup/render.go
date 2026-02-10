package setup

import "fmt"

func RenderConfig(in *SetupInput) []byte {
	return []byte(fmt.Sprintf(`
database:
  type: %s
  host: %s
  port: %d
  name: %s
  user: %s
  password: %s

backup:
  retention: %d
  dir: %s
`,
		in.DBType,
		in.Host,
		in.Port,
		in.DBName,
		in.User,
		in.Password,
		in.Retention,
		in.BackupDir,
	))
}
