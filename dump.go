package main

import (
	"fmt"
	"github.com/romanoff/dk/source"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var sourceRegistry map[string]Source = map[string]Source{
	"mysql": &source.Mysql{},
	"fs":    &source.Fs{},
	"mongo": &source.Mongo{},
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
		source := sourceRegistry[conf.Type]
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
	cmd := exec.Command("tar", "--remove-files", "-cjf", name+".tar.bz2", name)
	cmd.Dir = ".dklocal"
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ApplyDump(options map[string]string) {
	name := options["name"]
	if name == "" {
		ShowUsage()
		return
	}
	CheckConfig()
	UnarchiveBundle(name, func() {
		for sourceName, conf := range config.Sources {
			source := sourceRegistry[conf.Type]
			if source != nil {
				path := fmt.Sprintf(".dklocal/%v/%v", name, sourceName)
				// TODO: Check if this path exists
				source.Setup(&conf)
				err := source.ApplyDump(path)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}
	})
}

func UnarchiveBundle(name string, block func()) {
	cmd := exec.Command("tar", "-xf", name+".tar.bz2")
	cmd.Dir = ".dklocal"
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	block()
	err = os.RemoveAll(".dklocal/" + name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ListDumps(options map[string]string) {
	files, _ := ioutil.ReadDir(".dklocal/")
	for _, f := range files {
		bundle := f.Name()
		if strings.HasSuffix(bundle, ".tar.bz2") {
			fmt.Println(strings.Replace(bundle, ".tar.bz2", "", 1))
		}
	}
}
