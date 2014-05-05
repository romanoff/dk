package main

import (
	"fmt"
	"github.com/romanoff/dk/source"
	"os"
	"os/exec"
)

var sourceRegistry map[string]Source = map[string]Source{
	"mysql": &source.Mysql{},
}

func CreateDump(options map[string]string) {
	name := options["name"]
	if name == "" {
		ShowUsage()
		return
	}
	CheckConfig()
	perm := os.FileMode(0777)
	// TODO: Check if specified bundle has already been created
	for sourceName, conf := range config.Sources {
		source := sourceRegistry[sourceName]
		if source != nil {
			path := fmt.Sprintf(".dklocal/%v/%v", name, sourceName)
			os.MkdirAll(path, perm)
			source.Setup(&conf)
			err := source.CreateDump(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
	// TODO: Create timestamp (when bundle has been created) in case some sources were specified
	ArchiveBundle(name)
}

func ArchiveBundle(name string) {
	cmd := exec.Command("tar", "--remove-files", "-cjf", ".dklocal/"+name+".tar.bz2", ".dklocal/"+name)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
