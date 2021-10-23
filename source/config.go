package source

type Config struct {
	Type     string
	Name     string
	Password string
	Host     string
	Database string
	Port     string
	Paths    []string
	Plain    bool
}
