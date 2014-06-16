package main

import (
	"fmt"
	"github.com/romanoff/dk/remote"
	"os"
)

func PushDump(options map[string]string) {
	remoteName := options["remote"]
	remote, err := GetRemote(remoteName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if remote == nil {
		fmt.Printf("No remote with '%v' name found\n", remoteName)
	}
}

func PullDump(options map[string]string) {
	remoteName := options["remote"]
	remote, err := GetRemote(remoteName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if remote == nil {
		fmt.Printf("No remote with '%v' name found\n", remoteName)
	}
}

func ListRemoteDumps(options map[string]string) {
	remoteName := options["remote"]
	remote, err := GetRemote(remoteName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if remote == nil {
		fmt.Printf("No remote with '%v' name found\n", remoteName)
	}
}

func GetRemote(name string) (remote.Remote, error) {
	return nil, nil
}
