package source

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"errors"
	"path"
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

func (self *Mysql) CreateDump(dumpDir string) error {
	args := []string{}

	if self.Name != "" {
		args = append(args, "-u", self.Name)
	}
	if self.Password != "" {
		args = append(args, fmt.Sprintf("-p%v", self.Password))
	}
	if self.Host != "" {
		args = append(args, "--host", self.Host)
	}
	if self.Database == "" {
		return errors.New("Missing database name")
	}
	args = append(args, self.Database)

	out, err := os.OpenFile(path.Join(dumpDir, "dump.sql"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	cmd := exec.Command("mysqldump", args...)
	cmd.Stdout = out

	return cmd.Run()
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
