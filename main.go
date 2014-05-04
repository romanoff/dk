package main

import (
	"github.com/romanoff/gofh"
	"os"
)

type Source interface {
	CreateDump(config map[string]string) (error)
	ApplyDump(config map[string]string) (error)
}

func main() {
	fh := gofh.Init()
	fh.HandleCommand("init :source", CreateConfig)
	fh.Parse(os.Args[1:])	
}
