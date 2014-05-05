package main

import (
	"github.com/romanoff/dk/source"
	"github.com/romanoff/gofh"
	"os"
)

type Source interface {
	Setup(*source.Config)
	CreateDump(string) error
	ApplyDump(path string) error
}

var config *Config

func main() {
	config, _ = ReadConfig()
	fh := gofh.Init()
	fh.HandleCommand("init :source", CreateConfig)
	fh.HandleCommand("create :name", CreateDump)
	fh.HandleCommand("apply :name", ApplyDump)
	fh.HandleCommand("list", ListDumps)
	fh.Parse(os.Args[1:])
}
