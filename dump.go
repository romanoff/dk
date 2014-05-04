package main

import (
	"fmt"
	"github.com/romanoff/dk/source"
	"os"
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
	for sourceName, conf := range config.Sources {
		source := sourceRegistry[sourceName]
		if source != nil {
			source.Setup(&conf)
			err := source.CreateDump(name)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
