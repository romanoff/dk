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
	emptySchemaCommand := fmt.Sprintf("-u %v -p%v --host %v -e", self.Name, self.Password, self.Host)
	query := fmt.Sprintf("drop schema if exists %v;create schema %v;", self.Database, self.Database)
	params := strings.Split(emptySchemaCommand, " ")
	params = append(params, query)
	clearCommand := exec.Command("mysql", params...)
	err := clearCommand.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sqlFile, err := os.Open(path + "/dump.sql")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	applyCommand := fmt.Sprintf("-u %v -p%v --host %v --database %v", self.Name, self.Password, self.Host, self.Database)
	values := strings.Split(applyCommand, " ")
	cmd := exec.Command("mysql", values...)
	cmd.Stdin = sqlFile
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}
