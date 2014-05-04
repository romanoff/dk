package source

type Mysql struct {
	User     string
	Password string
	Host     string
	Database string
}

func (self *Mysql) Setup(config *Config) {
	self.User = config.User
	self.Password = config.Password
	self.Host = config.Host
	self.Database = config.Database
}

func (self *Mysql) CreateDump(dumpName string) error {
	return nil
}

func (self *Mysql) ApplyDump(path string) error {
	return nil
}
