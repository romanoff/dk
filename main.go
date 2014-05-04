package main

import (
	"github.com/romanoff/gofh"
	"os"
)

type Source interface {
	CreateDump() error
	ApplyDump(path string) error
}

func main() {
	fh := gofh.Init()
	fh.HandleCommand("init :source", CreateConfig)
	fh.Parse(os.Args[1:])
}
