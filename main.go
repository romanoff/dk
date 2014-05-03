package main

import (
	"github.com/romanoff/gofh"
	"os"
)

func main() {
	fh := gofh.Init()
	fh.HandleCommand("init :dbname", CreateConfig)
	fh.Parse(os.Args[1:])	
}
