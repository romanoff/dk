package source

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type Mongo struct {
	Host     string
	Database string
}

func (self *Mongo) Setup(config *Config) {
	self.Host = config.Host
	self.Database = config.Database
}

func (self *Mongo) CreateDump(dumpDir string) error {
	args := []string{}
	if self.Host != "" {
		args = append(args, "--host", self.Host)
	}
	if self.Database == "" {
		return errors.New("Missing database name")
	} else {
		args = append(args, "--db", self.Database)
	}
	args = append(args, "-o", dumpDir)
	cmd := exec.Command("mongodump", args...)
	return cmd.Run()
}

func (self *Mongo) ApplyDump(path string) error {
	args := []string{}

	if self.Host != "" {
		args = append(args, "--host", self.Host)
	}
	if self.Database == "" {
		return errors.New("Missing database name")
	} else {
		args = append(args, "--db", self.Database)
	}
	args = append(args, "--drop")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New(fmt.Sprintf("Expected to get one folder in mongo dump, but got %v", len(files)))
	}
	if !files[0].IsDir() {
		return errors.New(fmt.Sprintf("Expected to have one folder in mongo dump directory, but got file"))
	}
	dirPath := path + "/" + files[0].Name()
	args = append(args, dirPath)
	cmd := exec.Command("mongorestore", args...)
	return cmd.Run()
}
