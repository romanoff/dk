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
	config, _ = ReadConfig(".")
	fh := gofh.Init()
	fh.HandleCommand("init :source", CreateConfig)
	fh.HandleCommand("create :name", CreateDump)
	fh.HandleCommand("apply :name", ApplyDump)
	fh.HandleCommand("rm :name", RemoveDump)
	fh.HandleCommand("list", ListDumps)
	fh.HandleCommand("list :remote", ListRemoteDumps)
	fh.HandleCommand("push :remote :name", PushDump)
	fh.HandleCommand("pull :remote :name", PullDump)
	fh.HandleCommand("rm :remote :name", RemoveRemoteDump)
	fh.Parse(os.Args[1:])
}
