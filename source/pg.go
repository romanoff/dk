package source

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
)

type Postgres struct {
	Name     string
	Password string
	Host     string
	Database string
	Port     string
}

func (self *Postgres) Setup(config *Config) {
	self.Name = config.Name
	self.Password = config.Password
	self.Host = config.Host
	self.Port = config.Port
	self.Database = config.Database
}

func (self *Postgres) CreateDump(dumpDir string) error {
	args := []string{}
	if self.Host != "" {
		args = append(args, "-h", self.Host)
	} else {
		args = append(args, "-h", "127.0.0.1")
	}
	if self.Port != "" {
		args = append(args, "-p", self.Port)
	}
	args = append(args, "-Fc")
	if self.Name != "" {
		args = append(args, "-U", self.Name)
	}
	if self.Password != "" {
		os.Setenv("PGPASSWORD", self.Password)
	}
	if self.Database == "" {
		return errors.New("Missing database name")
	}
	args = append(args, self.Database)
	out, err := os.OpenFile(path.Join(dumpDir, "pg.dump"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	cmd := exec.Command("pg_dump", args...)
	cmd.Stdout = out

	return cmd.Run()
}

func (self *Postgres) ApplyDump(path string) error {
	args := []string{}
	if self.Host != "" {
		args = append(args, "-h", self.Host)
	} else {
		args = append(args, "-h", "127.0.0.1")
	}
	if self.Port != "" {
		args = append(args, "-p", self.Port)
	}
	args = append(args, "-Fc", "--clean")
	if self.Name != "" {
		args = append(args, "-U", self.Name)
	}
	if self.Password != "" {
		os.Setenv("PGPASSWORD", self.Password)
	}
	if self.Database == "" {
		return errors.New("Missing database name")
	}
	args = append(args, "-d", self.Database)
	args = append(args, path+"/pg.dump")
	cmd := exec.Command("pg_restore", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}
