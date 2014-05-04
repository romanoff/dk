package source

type Mysql struct {
	User     string
	Password string
	Host     string
	Database string
}

func (self *Mysql) CreateDump(config map[string]string) error {
	return nil
}

func (self *Mysql) ApplyDump(path string) error {
	return nil
}
