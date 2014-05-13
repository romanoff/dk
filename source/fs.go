package source

type Fs struct {
	Paths []string
}

func (self *Fs) Setup(config *Config) {
	self.Paths = config.Paths
}

func (self *Fs) CreateDump(path string) error {
	return nil
}

func (self *Fs) ApplyDump(path string) error {
	return nil
}
