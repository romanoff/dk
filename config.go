package main

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

func CreateConfig(options map[string]string) {
	perm := os.FileMode(0664)
	ioutil.WriteFile(".dk", []byte{}, perm)
	fmt.Println("dk config has been created")
}

type Config struct {
	Sources map[string]source
}

type source struct {
	Name     string
	Password string
	Host     string
	Database string
}

func ReadConfig() (*Config, error) {
	conf := &Config{}
	content, err := ioutil.ReadFile(".dk")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while reading .dk configuration: %v", err))
	}
	_, err = toml.Decode(string(content), &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
