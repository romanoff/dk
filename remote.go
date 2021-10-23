package main

import (
	"errors"
	"fmt"
	"github.com/romanoff/dk/remote"
	"os"
	"strings"
)

func PushDump(options map[string]string) {
	remoteName := options["remote"]
	r, err := GetRemote(remoteName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if r == nil {
		fmt.Printf("No remote with '%v' name found\n", remoteName)
	}
	dumpName := options["name"]
	dumpPath := "./.dklocal/" + dumpName + ".tar.bz2"
	// Check if dump with specified name exists
	err = r.Push(dumpPath, dumpName+".tar.bz2")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func PullDump(options map[string]string) {
	remoteName := options["remote"]
	r, err := GetRemote(remoteName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if r == nil {
		fmt.Printf("No remote with '%v' name found\n", remoteName)
	}
	dumpName := options["name"]
	dumpPath := "./.dklocal/" + dumpName + ".tar.bz2"
	err = r.Pull(dumpName+".tar.bz2", dumpPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ListRemoteDumps(options map[string]string) {
	remoteName := options["remote"]
	r, err := GetRemote(remoteName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if r == nil {
		fmt.Printf("No remote with '%v' name found\n", remoteName)
	}
	files, err := r.FilesList()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, bundle := range files {
		if strings.HasSuffix(bundle, ".tar.bz2") {
			fmt.Println(strings.Replace(bundle, ".tar.bz2", "", 1))
		}
	}
}

func RemoveRemoteDump(options map[string]string) {
	remoteName := options["remote"]
	r, err := GetRemote(remoteName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if r == nil {
		fmt.Printf("No remote with '%v' name found\n", remoteName)
	}
	dumpName := options["name"] + ".tar.bz2"
	err = r.Remove(dumpName)
	if err != nil {
		fmt.Println(err)
	}
}

var remoteRegistry map[string]remote.Remote = map[string]remote.Remote{
	"fs":  &remote.Fs{},
	"s3":  &remote.S3{},
	"scp": &remote.Scp{},
}

func GetRemote(name string) (remote.Remote, error) {
	remoteConfig, ok := config.Remotes[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("No remote '%v' found", name))
	}
	r := remoteRegistry[remoteConfig.Type]
	if r == nil {
		return nil, errors.New(fmt.Sprintf("No remote type '%v' found", remoteConfig.Type))
	}
	r.Setup(&remoteConfig)
	return r, nil
}
