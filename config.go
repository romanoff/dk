package main

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/romanoff/dk/remote"
	"github.com/romanoff/dk/source"
	"io/ioutil"
	"os"
)

func CreateConfig(options map[string]string) {
	perm := os.FileMode(0664)
	ioutil.WriteFile(".dk", []byte{}, perm)
	fmt.Println("dk config has been created")
}

type Config struct {
	Sources map[string]source.Config
	Remotes map[string]remote.Config
}

func ReadConfig(path string) (*Config, error) {
	conf := &Config{}
	content, err := ioutil.ReadFile(path + "/.dk")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while reading .dk configuration: %v", err))
	}
	_, err = toml.Decode(string(content), &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func CheckConfig() {
	if config == nil {
		fmt.Println(".dk config is missing")
		os.Exit(1)
	}
}
