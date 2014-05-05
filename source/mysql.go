package source

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Mysql struct {
	Name     string
	Password string
	Host     string
	Database string
}

func (self *Mysql) Setup(config *Config) {
	self.Name = config.Name
	self.Password = config.Password
	self.Host = config.Host
	self.Database = config.Database
}

func (self *Mysql) CreateDump(path string) error {
	dumpCommand := fmt.Sprintf("-u %v -p%v --host %v %v", self.Name, self.Password, self.Host, self.Database)
	values := strings.Split(dumpCommand, " ")
	cmd := exec.Command("mysqldump", values...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	perm := os.FileMode(0644)
	ioutil.WriteFile(path+"/dump.sql", output, perm)
	return nil
}

func (self *Mysql) ApplyDump(path string) error {
	return nil
}
