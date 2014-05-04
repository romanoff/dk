package main

import (
	"github.com/romanoff/gofh"
	"os"
)

type Source interface {
	CreateDump() error
	ApplyDump(path string) error
}

var config *Config

func main() {
	config, _ = ReadConfig()
	fh := gofh.Init()
	fh.HandleCommand("init :source", CreateConfig)
	fh.HandleCommand("create :name", CreateDump)
	fh.Parse(os.Args[1:])
}
